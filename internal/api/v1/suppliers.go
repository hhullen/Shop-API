package api

import (
	"net/http"
	ds "shopapi/internal/datastruct"
)

const (
	prefixSupplier        = apiPrefix + "/supplier"
	prefixSupplierAddress = prefixSupplier + "/address"
	prefixSuppliers       = apiPrefix + "/suppliers"
)

func (a *API) setupSuppliersHandlers(router IRouter) {
	router.HandleFunc(pattern(http.MethodPost, prefixSupplier), a.PutSupplier)
	router.HandleFunc(pattern(http.MethodPatch, prefixSupplierAddress), a.UpdateSupplierAddress)
	router.HandleFunc(pattern(http.MethodDelete, prefixSupplier), a.DeleteSupplier)
	router.HandleFunc(pattern(http.MethodGet, prefixSuppliers), a.GetSuppliers)
	router.HandleFunc(pattern(http.MethodGet, prefixSupplier), a.GetSupplier)
}

// PutSupplier Добавляет нового поставщика
// @Summary      Добавление поставщика
// @Description  Добавление поставщика. Если поставщик существует, вернется uid этого поставщика.
// @Tags         Supplier
// @Accept       json
// @Produce      json
// @Param        input body      ds.AddSupplierRequest  true "Информация о поставщике"
// @Success      200   {object}  ds.AddSupplierResponse
// @Failure      400   {object}  ds.AddSupplierResponse
// @Failure      500   {object}  ds.AddSupplierResponse
// @Router       /supplier [post]
func (a *API) PutSupplier(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.AddSupplierRequest, ds.AddSupplierResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.supplierService.AddSupplier,
	})
}

// UpdateSupplierAddress Обновляет адрес поставщика
// @Summary      Обновление адреса поставщика
// @Description  Обновление адреса поставщика
// @Tags         Supplier
// @Accept       json
// @Produce      json
// @Param        input body      ds.UpdateSupplierAddressRequest  true "uid и адрес"
// @Success      200   {object}  ds.UpdateSupplierAddressResponse
// @Failure      400   {object}  ds.UpdateSupplierAddressResponse
// @Failure      500   {object}  ds.UpdateSupplierAddressResponse
// @Router       /supplier/address [patch]
func (a *API) UpdateSupplierAddress(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.UpdateSupplierAddressRequest, ds.UpdateSupplierAddressResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.supplierService.UpdateSupplierAddress,
	})
}

// DeleteSupplier Удаляет поставщика
// @Summary      Удаление поставщика
// @Description  Удаление поставщика
// @Tags         Supplier
// @Accept       json
// @Produce      json
// @Param        input body      ds.DeleteSupplierRequest  true "uid"
// @Success      200   {object}  ds.DeleteSupplierResponse
// @Failure      400   {object}  ds.DeleteSupplierResponse
// @Failure      500   {object}  ds.DeleteSupplierResponse
// @Router       /supplier [delete]
func (a *API) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.DeleteSupplierRequest, ds.DeleteSupplierResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.supplierService.DeleteSupplier,
	})
}

// GetSupplier возвращает поставщика
// @Summary      Возвращает поставщика
// @Description  Возвращает поставщика
// @Tags         Supplier
// @Produce      json
// @Param        uid            query  string  true  "uid" example("609ccf6f-7fb4-44bd-aa77-bc9e0e7572b4")
// @Param        avoid_cache    query  string  false "avoid_cache" example(true)
// @Success      200  {object}  ds.GetSupplierResponse
// @Failure      400  {object}  ds.GetSupplierResponse
// @Failure      500  {object}  ds.GetSupplierResponse
// @Router       /supplier [get]
func (a *API) GetSupplier(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetSupplierRequest, ds.GetSupplierResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.supplierService.GetSupplier,
	})
}

// GetSuppliers возвращает поставщиков
// @Summary      Возвращает поставщиков
// @Description  Возвращает поставщиков
// @Tags         Supplier
// @Produce      json
// @Param        offset query     string                  true  "offset" example(0)
// @Param        limit  query     string                  true  "limit" example(10)
// @Success      200    {object}  ds.GetSuppliersResponse
// @Failure      400    {object}  ds.GetSuppliersResponse
// @Failure      500    {object}  ds.GetSuppliersResponse
// @Router       /suppliers [get]
func (a *API) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetSuppliersRequest, ds.GetSuppliersResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.supplierService.GetSuppliers,
	})
}
