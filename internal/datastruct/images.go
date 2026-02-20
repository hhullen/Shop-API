package datastruct

import "github.com/google/uuid"

type AddImageRequest struct {
	Uid   uuid.UUID `schema:"uid" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	Image []byte    `file:"image" validate:"required"`
}

type AddImageResponse struct {
	Status
	Uid *uuid.UUID `json:"uid,omitempty"`
}

type UpdateImageRequest struct {
	Uid   uuid.UUID `schema:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	Image []byte    `file:"image" validate:"required"`
}

type UpdateImageResponse struct {
	Status
}

type DeleteImageRequest struct {
	Uid uuid.UUID `json:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
}

type DeleteImageResponse struct {
	Status
}

type GetProductImageRequest struct {
	ProductUid uuid.UUID `schema:"product_uid" validate:"required" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
	AvoidCache bool      `schema:"avoid_cache,omitempty" example:"true"`
}

type GetProductImageResponse struct {
	Status
	Uid    *uuid.UUID `asFileName:"true" json:"uid,omitempty"`
	Image  []byte     `file:"image" json:"image,omitempty"`
	Cached bool       `json:"cached,omitempty" example:"false"`
}

type GetImageRequest struct {
	Uid        uuid.UUID `schema:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	AvoidCache bool      `schema:"avoid_cache,omitempty" example:"true"`
}

type GetImageResponse struct {
	Status
	Uid    *uuid.UUID `asFileName:"true" json:"uid,omitempty"`
	Image  []byte     `file:"image" json:"image,omitempty"`
	Cached bool       `json:"cached,omitempty" example:"false"`
}
