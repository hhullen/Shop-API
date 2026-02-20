package datastruct

import (
	"fmt"
	"shopapi/internal/supports"
	"strings"

	"github.com/google/uuid"
)

type PhoneNumber string

func (pn *PhoneNumber) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if err := supports.ValidatePhoneNumber(s); err != nil {
		return err
	}

	*pn = PhoneNumber(s)

	return nil
}

func (pn *PhoneNumber) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", *pn), nil
}

type Supplier struct {
	Uid         uuid.UUID   `json:"uid" example:"609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4"`
	PhoneNumber PhoneNumber `json:"phone_number" validate:"required" example:"+79336579933 RU"`
	Name        string      `json:"name" validate:"required" example:"Vasilisa&Drozzhi .ltd"`
	Address     *Address    `json:"address" validate:"required"`
}

type AddSupplierRequest struct {
	Supplier
}

type AddSupplierResponse struct {
	Uid *uuid.UUID `json:"uid,omitempty"`
}

type UpdateSupplierAddressRequest struct {
	Uid     uuid.UUID `json:"uid" validate:"required"  example:"609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4"`
	Address *Address  `json:"address" validate:"required"`
}

type UpdateSupplierAddressResponse struct {
	Status
}

type DeleteSupplierRequest struct {
	Uid uuid.UUID `json:"uid" validate:"required" example:"609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4"`
}

type DeleteSupplierResponse struct {
	Status
}

type GetSuppliersRequest struct {
	Limit  int64 `json:"limit" example:"10"`
	Offset int64 `json:"offset" example:"0"`
}

type GetSuppliersResponse struct {
	Suppliers []Supplier `json:"suppliers"`
}

type GetSupplierRequest struct {
	Uid        uuid.UUID `schema:"uid" validate:"required" example:"609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4"`
	AvoidCache bool      `schema:"avoid_cache,omitempty" example:"true"`
}

type GetSupplierResponse struct {
	Status
	Supplier *Supplier `json:"supplier,omitempty"`
	Cached   bool      `json:"cached,omitempty" example:"false"`
}
