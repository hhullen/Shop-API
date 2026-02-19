package api

import (
	"net/http"
	ds "shopapi/internal/datastruct"
)

const (
	prefixProduct  = apiPrefix + "/product"
	prefixProducts = apiPrefix + "/products"
)

func (a *API) setupProductsHandlers(router IRouter) {
	router.HandleFunc(pattern(http.MethodPost, prefixProduct), a.PutProduct)
	router.HandleFunc(pattern(http.MethodPatch, prefixProduct), a.DecreaseProduct)
	router.HandleFunc(pattern(http.MethodGet, prefixProduct), a.GetProduct)
	router.HandleFunc(pattern(http.MethodGet, prefixProducts), a.GetProducts)
	router.HandleFunc(pattern(http.MethodDelete, prefixProduct), a.DeleteProduct)
}

// PutProduct Добавляет новый продукт
// @Summary      Добавление продукта
// @Description  Добавление продукта. Если продукт существует, обновится цена, количество в наличии, дата обновления и вернется uid этого продукта.
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input body      ds.AddProductRequest  true "Информация о продукте"
// @Success      200   {object}  ds.AddProductResponse
// @Failure      400   {object}  ds.AddProductResponse
// @Failure      500   {object}  ds.Status
// @Router       /product [post]
func (a *API) PutProduct(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.AddProductRequest, ds.AddProductResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.productService.AddProduct,
	})
}

// DecreaseProduct Убавляет количество продукта
// @Summary      Убавление количества продукта
// @Description  Убавление количества продукта.
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input body      ds.DecreaseProductsRequest  true "uid и уколичество"
// @Success      200   {object}  ds.DecreaseProductsResponse
// @Failure      400   {object}  ds.DecreaseProductsResponse
// @Failure      500   {object}  ds.Status
// @Router       /product [patch]
func (a *API) DecreaseProduct(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.DecreaseProductsRequest, ds.DecreaseProductsResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.productService.DecreaseProducts,
	})
}

// GetProduct возвращает продукт
// @Summary      Возвращает продукт
// @Description  Возвращает продукт.
// @Tags         Product
// @Produce      json
// @Param        uid  query     string                true  "uid" example("c85a189d-d173-42e2-8e00-54395234d93d")
// @Success      200  {object}  ds.GetProductResponse
// @Failure      400  {object}  ds.GetProductResponse
// @Failure      500  {object}  ds.Status
// @Router       /product [get]
func (a *API) GetProduct(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetProductRequest, ds.GetProductResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.productService.GetProduct,
	})
}

// GetProducts возвращает список продуктов
// @Summary      Возвращает список продуктов
// @Description  Возвращает список продуктов. Если offset limit равны 0, вернет список всех
// @Tags         Product
// @Produce      json
// @Param        offset query    string                 true  "offset" example(0)
// @Param        limit  query    string                 true  "limit" example(10)
// @Success      200    {object} ds.GetProductsResponse
// @Failure      400    {object} ds.Status
// @Failure      500    {object} ds.Status
// @Router       /products [get]
func (a *API) GetProducts(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetProductsRequest, ds.GetProductsResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.productService.GetProducts,
	})
}

// DeleteProduct Удаляет продукт
// @Summary      Удаление продукта
// @Description  Удаление продукта.
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input body      ds.DeleteProductRequest  true "uid"
// @Success      200   {object}  ds.DeleteProductResponse
// @Failure      400   {object}  ds.DeleteProductResponse
// @Failure      500   {object}  ds.Status
// @Router       /product [delete]
func (a *API) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.DeleteProductRequest, ds.DeleteProductResponse]{
		api:              a,
		httpRequest:      r,
		httpResponse:     &w,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.productService.DeleteProduct,
	})
}
