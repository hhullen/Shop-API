package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ds "shopapi/internal/datastruct"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestPutClient(t *testing.T) {
	t.Parallel()

	t.Run("PutClient 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Client{
			Uid:              uid,
			Birthday:         ds.DateOnlyFromString("10.12.2011"),
			RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
			Name:             "Vasilisa",
			Surname:          "Kadyk",
			Gender:           ds.Male,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(&clientStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPut, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.AddClientResponse{
			Uid: &uid,
		}

		a.clientMock.EXPECT().AddClient(&ds.AddClientRequest{
			Client: clientStruct,
		}).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.PutClient(a.responseWriter, testReq)
	})

	t.Run("PutClient 400 No address", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Client{
			Uid:              uid,
			Birthday:         ds.DateOnlyFromString("10.12.2011"),
			RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
			Name:             "Vasilisa",
			Surname:          "Kadyk",
			Gender:           ds.Male,
		}

		jsonBody, err := json.Marshal(&clientStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPut, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutClient(a.responseWriter, testReq)
	})

	t.Run("PutClient 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid, err := uuid.Parse("4988150e-1c82-490f-8c07-ee74ace2dd14")
		if err != nil {
			t.Fatal(err)
		}

		clientStruct := ds.Client{
			Uid:              uid,
			Birthday:         ds.DateOnlyFromString("10.12.2011"),
			RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
			Name:             "Vasilisa",
			Surname:          "Kadyk",
			Gender:           ds.Male,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(&clientStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPut, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.clientMock.EXPECT().AddClient(&ds.AddClientRequest{Client: clientStruct}).Return(nil)
		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutClient(a.responseWriter, testReq)
	})
}

func TestDeleteClient(t *testing.T) {
	t.Parallel()

	t.Run("DeleteClient 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteClientRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.clientMock.EXPECT().DeleteClient(reqStruct).Return(&ds.DeleteClientResponse{Status: ds.Status{Message: ds.StatusOK}})
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteClient(a.responseWriter, testReq)
	})

	t.Run("DeleteClient 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteClientRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.clientMock.EXPECT().DeleteClient(reqStruct).Return(&ds.DeleteClientResponse{Status: ds.Status{Message: ds.StatusNotFound}})
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteClient(a.responseWriter, testReq)
	})

	t.Run("DeleteClient 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteClientRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.clientMock.EXPECT().DeleteClient(reqStruct).Return(&ds.DeleteClientResponse{Status: ds.Status{Message: ds.StatusServiceError}})
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteClient(a.responseWriter, testReq)
	})
}

func TestGetClientsByName(t *testing.T) {
	t.Parallel()

	t.Run("GetClientsByName 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Client{
			Uid:              uid,
			Birthday:         ds.DateOnlyFromString("10.12.2011"),
			RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
			Name:             "Vasilisa",
			Surname:          "Kadyk",
			Gender:           ds.Male,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixClientsByName, nil)

		q := testReq.URL.Query()
		q.Add("client_name", clientStruct.Name)
		q.Add("client_surname", clientStruct.Surname)
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetClientsByNameResponse{
			Clients: []ds.Client{clientStruct},
		}

		a.clientMock.EXPECT().GetClientsByName(&ds.GetClientsByNameRequest{
			Name:    clientStruct.Name,
			Surname: clientStruct.Surname,
		}).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetClientsByName(a.responseWriter, testReq)
	})

	t.Run("GetClientsByName 400 No parameter", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		clientStruct := ds.Client{
			Surname: "Kadyk",
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixClientsByName, nil)

		q := testReq.URL.Query()
		q.Add("client_surname", clientStruct.Surname)
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetClientsByNameResponse{
			Clients: []ds.Client{},
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetClientsByName(a.responseWriter, testReq)
	})

	t.Run("GetClientsByName 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		clientStruct := ds.Client{
			Name:    "Vasilisa",
			Surname: "Kadyk",
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixClientsByName, nil)

		q := testReq.URL.Query()
		q.Add("client_name", clientStruct.Name)
		q.Add("client_surname", clientStruct.Surname)
		testReq.URL.RawQuery = q.Encode()

		a.clientMock.EXPECT().GetClientsByName(&ds.GetClientsByNameRequest{
			Name:    clientStruct.Name,
			Surname: clientStruct.Surname,
		}).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetClientsByName(a.responseWriter, testReq)
	})
}

func TestGetClients(t *testing.T) {
	t.Parallel()

	t.Run("GetClients 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		clientStruct := ds.Client{
			Uid:              uid,
			Birthday:         ds.DateOnlyFromString("10.12.2011"),
			RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
			Name:             "Vasilisa",
			Surname:          "Kadyk",
			Gender:           ds.Male,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		limit := 10
		offset := 1

		testReq := httptest.NewRequest(http.MethodGet, prefixClients, nil)

		q := testReq.URL.Query()
		q.Add("limit", fmt.Sprint(limit))
		q.Add("offset", fmt.Sprint(offset))
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetClientsResponse{
			Clients: []ds.Client{clientStruct},
		}

		a.clientMock.EXPECT().GetClients(&ds.GetClientsRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
		}).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetClients(a.responseWriter, testReq)
	})

	t.Run("GetClients 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		limit := 10
		offset := 1

		testReq := httptest.NewRequest(http.MethodGet, prefixClients, nil)

		q := testReq.URL.Query()
		q.Add("limit", fmt.Sprint(limit))
		q.Add("offset", fmt.Sprint(offset))
		testReq.URL.RawQuery = q.Encode()

		a.clientMock.EXPECT().GetClients(&ds.GetClientsRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
		}).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetClients(a.responseWriter, testReq)
	})
}

func TestPatchClientAddress(t *testing.T) {
	t.Parallel()

	t.Run("PatchClientAddress 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.PatchClientAddressRequest{
			Uid: uuid.New(),
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixClientAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")
		resp := &ds.PatchClientAddressResponse{Status: ds.Status{Message: ds.StatusOK}}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.clientMock.EXPECT().PatchClientAddress(reqStruct).Return(resp)

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.PatchClientAddress(a.responseWriter, testReq)
	})

	t.Run("PatchClientAddress 400 No uid", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.PatchClientAddressRequest{
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixClientAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PatchClientAddress(a.responseWriter, testReq)
	})

	t.Run("PatchClientAddress 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.PatchClientAddressRequest{
			Uid: uuid.New(),
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixClientAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.clientMock.EXPECT().PatchClientAddress(reqStruct).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PatchClientAddress(a.responseWriter, testReq)
	})
}
