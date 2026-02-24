package service

import (
	"context"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"strings"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service ILogger,ICache,IClientStorage,IProductStorage,ISupplierStorage,IImageStorage

type ILogger interface {
	InfoKV(message string, argsKV ...any)
	WarnKV(message string, argsKV ...any)
	ErrorKV(message string, argsKV ...any)
	FatalKV(message string, argsKV ...any)
	Infof(message string, args ...any)
	Warnf(message string, args ...any)
	Errorf(message string, args ...any)
	Fatalf(message string, args ...any)
}

type ICache interface {
	Read(key string, v any) (bool, error)
	Write(key string, v any) error
}

type ICachedState interface {
	SetCached(bool)
}

type IClientStorage interface {
	AddClient(*ds.AddClientRequest) (*ds.AddClientResponse, error)
	DeleteClient(*ds.DeleteClientRequest) (*ds.DeleteClientResponse, error)
	GetClientsByName(*ds.GetClientsByNameRequest) (*ds.GetClientsByNameResponse, error)
	GetClients(*ds.GetClientsRequest) (*ds.GetClientsResponse, error)
	PatchClientAddress(*ds.PatchClientAddressRequest) (*ds.PatchClientAddressResponse, error)
}

type IProductStorage interface {
	AddProduct(*ds.AddProductRequest) (*ds.AddProductResponse, error)
	DecreaseProducts(*ds.DecreaseProductsRequest) (*ds.DecreaseProductsResponse, error)
	GetProduct(*ds.GetProductRequest) (*ds.GetProductResponse, error)
	GetProducts(*ds.GetProductsRequest) (*ds.GetProductsResponse, error)
	DeleteProduct(*ds.DeleteProductRequest) (*ds.DeleteProductResponse, error)
}

type ISupplierStorage interface {
	AddSupplier(*ds.AddSupplierRequest) (*ds.AddSupplierResponse, error)
	UpdateSupplierAddress(*ds.UpdateSupplierAddressRequest) (*ds.UpdateSupplierAddressResponse, error)
	DeleteSupplier(*ds.DeleteSupplierRequest) (*ds.DeleteSupplierResponse, error)
	GetSuppliers(*ds.GetSuppliersRequest) (*ds.GetSuppliersResponse, error)
	GetSupplier(*ds.GetSupplierRequest) (*ds.GetSupplierResponse, error)
}

type IImageStorage interface {
	AddImage(*ds.AddImageRequest) (*ds.AddImageResponse, error)
	UpdateImage(*ds.UpdateImageRequest) (*ds.UpdateImageResponse, error)
	DeleteImage(*ds.DeleteImageRequest) (*ds.DeleteImageResponse, error)
	GetProductImage(*ds.GetProductImageRequest) (*ds.GetProductImageResponse, error)
	GetImage(*ds.GetImageRequest) (*ds.GetImageResponse, error)
}

type Service struct {
	ctx             context.Context
	logger          ILogger
	cache           ICache
	clientStorage   IClientStorage
	productStorage  IProductStorage
	supplierStorage ISupplierStorage
	imageStorage    IImageStorage
}

func NewService(ctx context.Context, l ILogger, c ICache,
	cs IClientStorage,
	ps IProductStorage,
	ss ISupplierStorage,
	is IImageStorage) *Service {
	return &Service{
		ctx:             ctx,
		logger:          l,
		cache:           c,
		clientStorage:   cs,
		productStorage:  ps,
		supplierStorage: ss,
		imageStorage:    is,
	}
}

func (s *Service) logHandlerStatus(handlerName, status string) {
	if status != "" {
		s.logger.InfoKV(supports.Concat(handlerName, " status"), "status", status)
	}
}

func makeCacheKey(vv ...string) string {
	length := 0
	for i := range vv {
		length += len(vv[i])
	}

	var b strings.Builder
	b.Grow(length)

	for i := range vv {
		b.WriteString(vv[i])
		b.WriteByte('_')
	}

	return b.String()
}

func execWithCache[RespT ICachedState](s *Service, key string, avoidCache bool, fetch func() (RespT, error)) (RespT, error) {
	var response RespT
	var cached bool
	var err error

	if !avoidCache {
		cached, err = s.cache.Read(key, &response)
		if err != nil {
			s.logger.ErrorKV("failed reading cache", "message", err.Error())
		}
	}

	if cached {
		response.SetCached(true)
		return response, nil
	}

	response, err = fetch()
	if err != nil {
		var empty RespT
		return empty, err
	}

	err = s.cache.Write(key, response)
	if err != nil {
		s.logger.ErrorKV("failed writing cache", "message", err.Error())
	}
	response.SetCached(false)

	return response, nil
}
