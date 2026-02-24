package api

import (
	"net/http"
	ds "shopapi/internal/datastruct"
)

const (
	prefixClient        = apiPrefix + "/client"
	prefixClientAddress = prefixClient + "/address"
	prefixClients       = apiPrefix + "/clients"
	prefixClientsByName = prefixClients + "/named"
)

func (a *API) setupClientsHandlers(router IRouter) {
	router.HandleFunc(pattern(http.MethodPost, prefixClient), a.PutClient)
	router.HandleFunc(pattern(http.MethodDelete, prefixClient), a.DeleteClient)
	router.HandleFunc(pattern(http.MethodGet, prefixClients), a.GetClients)
	router.HandleFunc(pattern(http.MethodGet, prefixClientsByName), a.GetClientsByName)
	router.HandleFunc(pattern(http.MethodPatch, prefixClientAddress), a.PatchClientAddress)
}

// PutClient Добавляет нового клиента
// @Summary      Добавление клиента
// @Description  Добавление клиента. Если клиент существует вернется uid существующего клиента.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        input body      ds.AddClientRequest  true "Информация о клиенте"
// @Success      200   {object}  ds.AddClientResponse
// @Failure      400   {object}  ds.Status
// @Failure      500   {object}  ds.Status
// @Router       /client [post]
func (a *API) PutClient(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.AddClientRequest, ds.AddClientResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.clientService.AddClient,
	})
}

// DeleteClient удаляет клиента
// @Summary      Удаление клиента
// @Description  Удаляет клиента по его uid
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        input body      ds.DeleteClientRequest  true "uid клиента"
// @Success      200   {object}  ds.DeleteClientResponse
// @Failure      400   {object}  ds.DeleteClientResponse
// @Failure      500   {object}  ds.Status
// @Router       /client [delete]
func (a *API) DeleteClient(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.DeleteClientRequest, ds.DeleteClientResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.clientService.DeleteClient,
	})
}

// GetClientsByName возвращает клиентов по имени и фамилии
// @Summary      Возвращает клиентов по имени и фамилии
// @Description  Возвращает клиентов по имени и фамилии
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        client_name    query  string false "client_name" example(Vasilisa)
// @Param        client_surname query  string false "client_surname" example(Kadyk)
// @Param        avoid_cache    query  string false "avoid_cache" example(true)
// @Success      200  {object}  ds.GetClientsByNameResponse
// @Failure      400  {object}  ds.Status
// @Failure      500  {object}  ds.Status
// @Router       /clients/named [get]
func (a *API) GetClientsByName(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetClientsByNameRequest, ds.GetClientsByNameResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.clientService.GetClientsByName,
	})
}

// GetClients возвращает список клиентов
// @Summary      Возвращает список клиентов
// @Description  Возвращает список клиентов. Если offset limit равны 0, вернет список всех клиентов
// @Tags         Client
// @Produce      json
// @Param        offset       query  string true  "offset"      example(0)
// @Param        limit        query  string true  "limit"       example(10)
// @Param        avoid_cache  query  string false "avoid_cache" example(true)
// @Success      200  {object}  ds.GetClientsResponse
// @Failure      400  {object}  ds.Status
// @Failure      500  {object}  ds.Status
// @Router       /clients [get]
func (a *API) GetClients(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.GetClientsRequest, ds.GetClientsResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractSchemaQuery,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.clientService.GetClients,
	})
}

// PatchClientAddress обновляет адрес клиента
// @Summary      Обновляет адрес клиента
// @Description  Обновляет адрес клиента по его uid
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        input body     ds.PatchClientAddressRequest  true "uid и адрес"
// @Success      200   {object} ds.PatchClientAddressResponse
// @Failure      400   {object} ds.PatchClientAddressResponse
// @Failure      500   {object} ds.Status
// @Router       /client/address [patch]
func (a *API) PatchClientAddress(w http.ResponseWriter, r *http.Request) {
	Exec(ExecArgs[ds.PatchClientAddressRequest, ds.PatchClientAddressResponse]{
		httpRequest:      r,
		httpResponse:     &w,
		api:              a,
		requestExtractor: extractJsonBody,
		responseWriter:   writeJsonResponse,
		serviceFunc:      a.clientService.PatchClientAddress,
	})
}
