package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"

	ds "shopapi/internal/datastruct"
	mimeManager "shopapi/internal/mime-manager"
	"shopapi/internal/service"
	"shopapi/internal/supports"

	_ "shopapi/internal/docs"

	"github.com/gorilla/schema"
	httpSwagger "github.com/swaggo/http-swagger"
)

//go:generate mockgen -source=api.go -destination=api_mock.go -package=api IClientService,IProductService,ISupplierService,IImageService,IWithStatus,IServer,IRouter
//go:generate mockgen -destination=http_mock.go -package=api net/http ResponseWriter

const (
	address                    = ":8080"
	readTimeout                = time.Second * 5
	writeTimeout               = time.Second * 5
	maxMultipartFormMemory10MB = 10 << 20

	contentTypeKey        = "Content-Type"
	contentLenKey         = "Content-Length"
	contentCachingKey     = "Cache-Control"
	contentDispositionKey = "Content-Disposition"

	appJSONValue         = "application/json"
	appOctetStream       = "application/octet-stream"
	appMiltipartFormData = "multipart/form-data"
	appPublicCacheOneDay = "public, max-age=86400"

	apiPrefix     = "/api/v1"
	swaggerPrefix = "/swagger/"
)

var schemaDecoder = schema.NewDecoder()

type IClientService interface {
	AddClient(*ds.AddClientRequest) *ds.AddClientResponse
	DeleteClient(*ds.DeleteClientRequest) *ds.DeleteClientResponse
	GetClientsByName(*ds.GetClientsByNameRequest) *ds.GetClientsByNameResponse
	GetClients(*ds.GetClientsRequest) *ds.GetClientsResponse
	PatchClientAddress(*ds.PatchClientAddressRequest) *ds.PatchClientAddressResponse
}

type IProductService interface {
	AddProduct(*ds.AddProductRequest) *ds.AddProductResponse
	DecreaseProducts(*ds.DecreaseProductsRequest) *ds.DecreaseProductsResponse
	GetProduct(*ds.GetProductRequest) *ds.GetProductResponse
	GetProducts(*ds.GetProductsRequest) *ds.GetProductsResponse
	DeleteProduct(*ds.DeleteProductRequest) *ds.DeleteProductResponse
}

type ISupplierService interface {
	AddSupplier(*ds.AddSupplierRequest) *ds.AddSupplierResponse
	UpdateSupplierAddress(*ds.UpdateSupplierAddressRequest) *ds.UpdateSupplierAddressResponse
	DeleteSupplier(*ds.DeleteSupplierRequest) *ds.DeleteSupplierResponse
	GetSuppliers(*ds.GetSuppliersRequest) *ds.GetSuppliersResponse
	GetSupplier(*ds.GetSupplierRequest) *ds.GetSupplierResponse
}

type IImageService interface {
	AddImage(*ds.AddImageRequest) *ds.AddImageResponse
	UpdateImage(*ds.UpdateImageRequest) *ds.UpdateImageResponse
	DeleteImage(*ds.DeleteImageRequest) *ds.DeleteImageResponse
	GetProductImage(*ds.GetProductImageRequest) *ds.GetProductImageResponse
	GetImage(*ds.GetImageRequest) *ds.GetImageResponse
}

type IWithStatus interface {
	GetStatus() string
}

type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type IRouter interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type API struct {
	ctx             context.Context
	server          IServer
	router          IRouter
	logger          service.ILogger
	clientService   IClientService
	productService  IProductService
	supplierService ISupplierService
	imageService    IImageService
}

type ExecArgs[ReqT any, RespT any] struct {
	api              *API
	serviceFunc      func(*ReqT) *RespT
	requestExtractor func(r *http.Request, v any) error
	responseWriter   func(r *http.ResponseWriter, v any) error
	httpRequest      *http.Request
	httpResponse     *http.ResponseWriter
}

var mutex sync.Mutex

var statusCodeMap = map[string]int{
	ds.StatusNotFound:     http.StatusNotFound,
	ds.StatusServiceError: http.StatusInternalServerError,
	ds.StatusOK:           http.StatusOK,
}

func getStatusCode(s string) (int, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	code, ok := statusCodeMap[s]
	return code, ok
}

func NewAPI(ctx context.Context, l service.ILogger,
	cs IClientService,
	ps IProductService,
	ss ISupplierService,
	is IImageService) *API {

	router := http.NewServeMux()
	router.Handle(swaggerPrefix, httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         address,
		Handler:      middlewareHandler(router),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	go func() {
		<-ctx.Done()
		err := server.Shutdown(context.Background())
		if err != nil {
			l.FatalKV("server failed graceful shutdown", "error", err.Error())
		}
	}()

	return buildAPI(ctx, l, server, router, cs, ps, ss, is)
}

func buildAPI(ctx context.Context,
	l service.ILogger,
	s IServer,
	r IRouter,
	cs IClientService,
	ps IProductService,
	ss ISupplierService,
	is IImageService) *API {
	api := &API{
		ctx:             ctx,
		server:          s,
		router:          r,
		clientService:   cs,
		productService:  ps,
		supplierService: ss,
		imageService:    is,
		logger:          l,
	}

	api.setupClientsHandlers(api.router)
	api.setupProductsHandlers(api.router)
	api.setupSuppliersHandlers(api.router)
	api.setupImagesHandlers(api.router)

	mimeManager.AddAllowedExtensions("image", []string{
		".jpg",
		".jpeg",
		".png",
		".webp",
	})

	return api
}

func (a *API) Start() error {
	a.logger.Infof("Server is listening")
	return a.server.ListenAndServe()
}

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func pattern(method, prefixPath string) string {
	return supports.Concat(method, " ", prefixPath)
}

func extractJsonBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func writeJsonResponse(w *http.ResponseWriter, resp any) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(resp); err != nil {
		return err
	}

	(*w).Header().Set(contentLenKey, strconv.Itoa(len(buf.Bytes())))
	(*w).Header().Set(contentTypeKey, appJSONValue)

	statusNum := http.StatusOK
	status, has := hasResponseStatus(resp)
	if has {
		if code, ok := getStatusCode(status); ok {
			statusNum = code
		} else if status != "" {
			statusNum = http.StatusBadRequest
		}
	}

	(*w).WriteHeader(statusNum)
	_, err := (*w).Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func hasResponseStatus(resp any) (string, bool) {
	if v, withStatus := resp.(IWithStatus); withStatus {
		return v.GetStatus(), true
	}

	return "", false
}

func extractSchemaQuery(r *http.Request, v any) error {
	if err := schemaDecoder.Decode(v, r.URL.Query()); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func extractMultipartWithFile(r *http.Request, v any) error {
	if err := r.ParseMultipartForm(maxMultipartFormMemory10MB); err != nil {
		return err
	}

	if err := schemaDecoder.Decode(v, r.PostForm); err != nil && err != io.EOF {
		return err
	}

	formName, field, err := supports.GetStructFieldByTagKey(v, "file")
	if err != nil {
		return err
	}

	if !supports.IsFieldByteSlice(field) {
		return fmt.Errorf("field '%s' has not '[]byte' type", formName)
	}

	file, _, err := r.FormFile(formName)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	err = mimeManager.IsFileAllowed(data, formName)
	if err != nil {
		return err
	}

	field.SetBytes(data)

	return nil
}

func writeFileResponse(w *http.ResponseWriter, resp any) error {
	value := reflect.ValueOf(resp)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct but got '%s'", value.Kind().String())
	}

	formName, field, err := supports.GetStructFieldByTagKey(resp, "file")

	if err != nil {
		return err
	}

	if !supports.IsFieldByteSlice(field) {
		return fmt.Errorf("field '%s' has not '[]byte' type", formName)
	}

	data := field.Bytes()

	if len(data) == 0 {
		return writeJsonResponse(w, resp)
	}

	filename, err := getAttachmentFileName(resp, data, formName)
	if err != nil {
		return err
	}

	(*w).Header().Set(contentDispositionKey, filename)
	(*w).Header().Set(contentLenKey, strconv.Itoa(len(data)))
	(*w).Header().Set(contentTypeKey, appOctetStream)
	// (*w).Header().Set(contentCachingKey, appPublicCacheOneDay)

	(*w).WriteHeader(http.StatusOK)
	_, err = (*w).Write(data)
	if err != nil {
		return err
	}

	return nil
}

func getAttachmentFileName(resp any, outData []byte, expectFileType string) (string, error) {
	var fieldName string
	_, field, err := supports.GetStructFieldByTagKey(resp, "asFileName")
	if err == nil {
		if v, ok := field.Interface().(fmt.Stringer); ok {
			fieldName = v.String()
		}
	}

	ext, err := mimeManager.GetFileExtension(outData, expectFileType)
	if err != nil {
		return "", err
	}
	return supports.Concat("attachment; filename=\"file_", fieldName, "_", supports.GetDateAsFileName(time.Now()), ext, "\""), nil
	// return fmt.Sprintf("attachment; filename=\"file_%s_%s%s\"",
	// 	fieldName, supports.GetDateAsFileName(time.Now()), ext), nil
}

func Exec[ReqT any, RespT any](a ExecArgs[ReqT, RespT]) {
	var req ReqT

	if err := a.requestExtractor(a.httpRequest, &req); err != nil {
		msg := "failed extracting request"
		a.api.logger.ErrorKV(msg, "error", err.Error())

		resp := ds.Status{Message: supports.Concat(msg, ": ", err.Error())}
		err = writeJsonResponse(a.httpResponse, resp)
		if err != nil {
			a.api.logger.ErrorKV("failed write response",
				"error", err.Error(), "response", resp)
		}

		return
	}

	if err := supports.StructValidator().Struct(&req); err != nil {
		msg := "failed validating request"
		a.api.logger.ErrorKV(msg, "error", err.Error(), "request", req)

		resp := ds.Status{Message: supports.Concat(msg, ": ", err.Error())}
		err = writeJsonResponse(a.httpResponse, resp)
		if err != nil {
			a.api.logger.ErrorKV("failed write response",
				"error", err.Error(), "response", resp)
		}
		return
	}

	resp := a.serviceFunc(&req)
	if resp == nil {
		resp := ds.Status{Message: ds.StatusServiceError}
		msg := "failed execute request on service"
		a.api.logger.ErrorKV(msg, "error", "service return no response", "request", req)
		err := writeJsonResponse(a.httpResponse, resp)
		if err != nil {

		}
		return
	}

	if err := a.responseWriter(a.httpResponse, resp); err != nil {
		msg := "failed writing response"
		http.Error(*a.httpResponse, msg, http.StatusInternalServerError)
		a.api.logger.ErrorKV(msg, "error", err.Error(), "request", req)
	}
}
