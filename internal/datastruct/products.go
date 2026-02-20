package datastruct

import (
	"github.com/google/uuid"
)

const (
	StatusDecreaseProductsFailed          = "not enough to decrease"
	StatusAddProductWithNoImageOrSupplier = "not exists image or supplier"
)

type Product struct {
	Uid             uuid.UUID `json:"uid" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
	SupplierUid     uuid.UUID `json:"supplier_id" validate:"required" example:"609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4"`
	ImageUid        uuid.UUID `json:"image_id" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	LastUpdateDate  DateOnly  `json:"last_update_date" example:"31.01.2026"`
	Name            string    `json:"name" validate:"required" example:"Wooden beam"`
	Category        string    `json:"category" validate:"required" example:"construction"`
	Price           float64   `json:"price" validate:"required" example:"299.95"`
	AvaliableStocks int64     `json:"available_stock" validate:"required" example:"1023"`
}

type AddProductRequest struct {
	Product
}

type AddProductResponse struct {
	Status
	Uid *uuid.UUID `json:"uid,omitempty"`
}

type DecreaseProductsRequest struct {
	Uid    uuid.UUID `json:"uid" validate:"required" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
	Amount int64     `json:"amount" validate:"required" example:"3"`
}

type DecreaseProductsResponse struct {
	Status
	Left *int64 `json:"left,omitempty"`
}

type GetProductRequest struct {
	Uid        uuid.UUID `schema:"uid" validate:"required" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
	AvoidCache bool      `schema:"avoid_cache,omitempty" example:"true"`
}

type GetProductResponse struct {
	Status
	Product *Product `json:"product,omitempty"`
	Cached  bool     `json:"cached,omitempty" example:"false"`
}

type GetProductsRequest struct {
	Limit  int64 `schema:"limit" example:"10"`
	Offset int64 `schema:"offset" example:"0"`
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
}

type DeleteProductRequest struct {
	Uid uuid.UUID `json:"uid" validate:"required" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
}

type DeleteProductResponse struct {
	Status
}
