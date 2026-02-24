package datastruct

import "github.com/google/uuid"

type AddImageRequest struct {
	AvoidCacheFlag
	Uid   uuid.UUID `schema:"uid" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	Image []byte    `file:"image" validate:"required"`
}

type AddImageResponse struct {
	Status
	CachedStatus
	Uid *uuid.UUID `json:"uid,omitempty"`
}

type UpdateImageRequest struct {
	AvoidCacheFlag
	Uid   uuid.UUID `schema:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
	Image []byte    `file:"image" validate:"required"`
}

type UpdateImageResponse struct {
	Status
	CachedStatus
}

type DeleteImageRequest struct {
	AvoidCacheFlag
	Uid uuid.UUID `json:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
}

type DeleteImageResponse struct {
	Status
	CachedStatus
}

type GetProductImageRequest struct {
	AvoidCacheFlag
	ProductUid uuid.UUID `schema:"product_uid" validate:"required" example:"c85a189d-d173-42e2-8e00-54395234d93d"`
}

type GetProductImageResponse struct {
	Status
	CachedStatus
	Uid   *uuid.UUID `asFileName:"true" json:"uid,omitempty"`
	Image []byte     `file:"image" json:"image,omitempty"`
}

type GetImageRequest struct {
	AvoidCacheFlag
	Uid uuid.UUID `schema:"uid" validate:"required" example:"376de312-5bcb-4320-8ba3-bd2050548229"`
}

type GetImageResponse struct {
	Status
	CachedStatus
	Uid   *uuid.UUID `asFileName:"true" json:"uid,omitempty"`
	Image []byte     `file:"image" json:"image,omitempty"`
}
