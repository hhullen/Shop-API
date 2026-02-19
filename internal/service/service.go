package service

import (
	"context"
	"fmt"
	ds "shopapi/internal/datastruct"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service ILogger,IClientStorage,IProductStorage,ISupplierStorage,IImageStorage

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
	clientStorage   IClientStorage
	productStorage  IProductStorage
	supplierStorage ISupplierStorage
	imageStorage    IImageStorage
}

func NewService(ctx context.Context, l ILogger,
	cs IClientStorage,
	ps IProductStorage,
	ss ISupplierStorage,
	is IImageStorage) *Service {
	return &Service{
		ctx:             ctx,
		logger:          l,
		clientStorage:   cs,
		productStorage:  ps,
		supplierStorage: ss,
		imageStorage:    is,
	}
}

func (s *Service) logHandlerStatus(handlerName, status string) {
	if status != "" {
		s.logger.InfoKV(fmt.Sprintf("%s status", handlerName), "status", status)
	}
}
