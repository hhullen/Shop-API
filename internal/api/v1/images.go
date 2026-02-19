package api

import (
	"net/http"
	ds "shopapi/internal/datastruct"
)

const (
	prefixImage        = apiPrefix + "/image"
	prefixImageProduct = prefixImage + "/product"
)

func (a *API) setupImagesHandlers(router IRouter) {
	router.HandleFunc(pattern(http.MethodPost, prefixImage), a.PutImage)
	router.HandleFunc(pattern(http.MethodPatch, prefixImage), a.UpdateImage)
	router.HandleFunc(pattern(http.MethodGet, prefixImageProduct), a.GetProductImage)
	router.HandleFunc(pattern(http.MethodGet, prefixImage), a.GetImage)
	router.HandleFunc(pattern(http.MethodDelete, prefixImage), a.DeleteImage)
}

// PutImage добавляет новое изображение
// @Summary      Добавляет новое изображение
// @Description  добавляет новое изображение
// @Tags         Image
// @Accept       mpfd
// @Produce      json
// @Param        uid   formData  string              false "uid" example("376de312-5bcb-4320-8ba3-bd2050548229")
// @Param        image formData  file                true  "Файл изображения"
// @Success      200   {object}  ds.AddImageResponse
// @Failure      400   {object}  ds.AddImageResponse
// @Failure      500   {object}  ds.Status
// @Router       /image [post]
func (a *API) PutImage(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.AddImageRequest, ds.AddImageResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractMultipartWithFile,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.imageService.AddImage,
	})
}

// UpdateImage обновление изображение
// @Summary      обновить изображение
// @Description  обновить существующее изображение
// @Tags         Image
// @Accept       mpfd
// @Produce      json
// @Param        uid   formData  string                 true "uid" example("376de312-5bcb-4320-8ba3-bd2050548229")
// @Param        image formData  file                   true "Файл изображения"
// @Success      200   {object}  ds.UpdateImageResponse
// @Failure      400   {object}  ds.UpdateImageResponse
// @Failure      500   {object}  ds.Status
// @Router       /image [patch]
func (a *API) UpdateImage(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.UpdateImageRequest, ds.UpdateImageResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractMultipartWithFile,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.imageService.UpdateImage,
	})
}

// GetProductImage возвращает изображение продукта
// @Summary      Возвращает изображение продукта
// @Description  Возвращает изображение продукта
// @Tags         Image
// @Produce      application/octet-stream
// @Param        product_uid    query  string  true  "product_uid" example("c85a189d-d173-42e2-8e00-54395234d93d")
// @Success      200  {file}    binary
// @Failure      400  {object}  ds.GetProductImageResponse
// @Failure      500  {object}  ds.Status
// @Router       /image/product [get]
func (a *API) GetProductImage(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetProductImageRequest, ds.GetProductImageResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeFileResponse,
		serviceFunc:      a.imageService.GetProductImage,
	})
}

// GetImage возвращает изображение
// @Summary      Возвращает изображение
// @Description  Возвращает изображение
// @Tags         Image
// @Produce      application/octet-stream
// @Param        uid  query     string              true  "uid" example("376de312-5bcb-4320-8ba3-bd2050548229")
// @Success      200  {file}    binary
// @Failure      400  {object}  ds.GetImageResponse
// @Failure      500  {object}  ds.Status
// @Router       /image [get]
func (a *API) GetImage(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetImageRequest, ds.GetImageResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeFileResponse,
		serviceFunc:      a.imageService.GetImage,
	})
}

// DeleteImage удаляет изображение
// @Summary      удаляет изображение
// @Description  удаляет изображение
// @Tags         Image
// @Accept       json
// @Produce      json
// @Param        input body      ds.DeleteImageRequest  true "uid"
// @Success      200   {object}  ds.DeleteImageResponse
// @Failure      400   {object}  ds.DeleteImageResponse
// @Failure      500   {object}  ds.Status
// @Router       /image [delete]
func (a *API) DeleteImage(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.DeleteImageRequest, ds.DeleteImageResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.imageService.DeleteImage,
	})
}
