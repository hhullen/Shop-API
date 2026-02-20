package datastruct

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g *Gender) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == string(Male) || s == string(Female) {
		*g = Gender(s)
		return nil
	}

	return fmt.Errorf("incorrect gender: '%s'", s)
}

func (g *Gender) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", *g), nil
}

type Client struct {
	Uid              uuid.UUID `json:"uid" example:"4988150e-1c82-490f-8c07-ee74ace2dd14"`
	Birthday         DateOnly  `json:"birthday" validate:"required" example:"10.12.2011"`
	RegistrationDate DateOnly  `json:"registration_date" validate:"required" example:"30/01/2026"`
	Name             string    `json:"client_name" validate:"required" example:"Vasilisa"`
	Surname          string    `json:"client_surname" validate:"required" example:"Kadyk"`
	Gender           Gender    `json:"gender" validate:"required" example:"female"`
	Address          *Address  `json:"address" validate:"required"`
}

type AddClientRequest struct {
	Client
}

type AddClientResponse struct {
	Uid *uuid.UUID `json:"uid,omitempty"`
}

type Address struct {
	Country string `json:"country" validate:"required" example:"USA"`
	City    string `json:"city" validate:"required" example:"Seattle"`
	Street  string `json:"street" validate:"required" example:"12th Ave E"`
}

type DeleteClientRequest struct {
	Uid uuid.UUID `json:"uid" validate:"required" example:"4988150e-1c82-490f-8c07-ee74ace2dd14"`
}

type DeleteClientResponse struct {
	Status
}

type GetClientsByNameRequest struct {
	Name       string `schema:"client_name" validate:"required" example:"Vasilisa"`
	Surname    string `schema:"client_surname" validate:"required" example:"Kadyk"`
	AvoidCache bool   `schema:"avoid_cache,omitempty" example:"true"`
}

type GetClientsByNameResponse struct {
	Clients []Client `json:"clients"`
	Cached  bool     `json:"cached,omitempty" example:"false"`
}

type GetClientsRequest struct {
	Limit  int64 `schema:"limit" example:"10"`
	Offset int64 `schema:"offset" example:"0"`
}

type GetClientsResponse struct {
	Clients []Client `json:"clients"`
}

type PatchClientAddressRequest struct {
	Uid     uuid.UUID `json:"uid" validate:"required" example:"4988150e-1c82-490f-8c07-ee74ace2dd14"`
	Address *Address  `json:"address" validate:"required"`
}

type PatchClientAddressResponse struct {
	Status
}
