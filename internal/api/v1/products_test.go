package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	ds "shopapi/internal/datastruct"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestPutProduct(t *testing.T) {
	t.Parallel()

	t.Run("PutProduct 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		reqStruct := ds.Product{
			Uid:             uuid.New(),
			SupplierUid:     uuid.New(),
			ImageUid:        uuid.New(),
			LastUpdateDate:  ds.DateOnlyFromString("01.01.2026"),
			Name:            "name",
			Category:        "category",
			Price:           299.99,
			AvaliableStocks: 20,
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPut, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		req := &ds.AddProductRequest{
			Product: reqStruct,
		}

		resp := &ds.AddProductResponse{
			Uid: &uid,
		}

		a.productMock.EXPECT().AddProduct(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.PutProduct(a.responseWriter, apiReq)
	})

	t.Run("PutProduct 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := ds.Product{
			Uid: uuid.New(),
			// SupplierUid:     uuid.New(),
			ImageUid:        uuid.New(),
			LastUpdateDate:  ds.DateOnlyFromString("01.01.2026"),
			Name:            "name",
			Category:        "category",
			Price:           299.99,
			AvaliableStocks: 20,
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPut, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutProduct(a.responseWriter, apiReq)
	})

	t.Run("PutProduct 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := ds.Product{
			Uid:             uuid.New(),
			SupplierUid:     uuid.New(),
			ImageUid:        uuid.New(),
			LastUpdateDate:  ds.DateOnlyFromString("01.01.2026"),
			Name:            "name",
			Category:        "category",
			Price:           299.99,
			AvaliableStocks: 20,
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPut, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		req := &ds.AddProductRequest{
			Product: reqStruct,
		}

		a.productMock.EXPECT().AddProduct(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutProduct(a.responseWriter, apiReq)
	})
}

func TestDecreaseProduct(t *testing.T) {
	t.Parallel()

	t.Run("DecreaseProduct 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: int64(12),
		}

		jsonBody, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPatch, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		left := int64(20)
		resp := &ds.DecreaseProductsResponse{
			Left: &left,
		}

		a.productMock.EXPECT().DecreaseProducts(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DecreaseProduct(a.responseWriter, apiReq)
	})

	t.Run("DecreaseProduct 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DecreaseProductsRequest{
			Uid: uid,
			// Amount: int64(12),
		}

		jsonBody, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPatch, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DecreaseProduct(a.responseWriter, apiReq)
	})

	t.Run("DecreaseProduct 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: int64(12),
		}

		jsonBody, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPatch, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		resp := &ds.DecreaseProductsResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}

		a.productMock.EXPECT().DecreaseProducts(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DecreaseProduct(a.responseWriter, apiReq)
	})

	t.Run("DecreaseProduct 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: int64(12),
		}

		jsonBody, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodPatch, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		a.productMock.EXPECT().DecreaseProducts(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DecreaseProduct(a.responseWriter, apiReq)
	})
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	t.Run("GetProduct 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.GetProductRequest{
			Uid: uid,
		}

		apiReq := httptest.NewRequest(http.MethodGet, prefixProduct, nil)
		q := apiReq.URL.Query()
		q.Add("uid", uid.String())
		apiReq.URL.RawQuery = q.Encode()

		resp := &ds.GetProductResponse{
			Product: &ds.Product{
				Uid:             uuid.New(),
				SupplierUid:     uuid.New(),
				ImageUid:        uuid.New(),
				LastUpdateDate:  ds.DateOnlyFromString("01.01.2026"),
				Name:            "name",
				Category:        "category",
				Price:           299.99,
				AvaliableStocks: 20,
			},
		}

		a.productMock.EXPECT().GetProduct(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetProduct(a.responseWriter, apiReq)
	})

	t.Run("GetProduct 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		apiReq := httptest.NewRequest(http.MethodGet, prefixProduct, nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetProduct(a.responseWriter, apiReq)
	})

	t.Run("GetProduct 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.GetProductRequest{
			Uid: uid,
		}

		apiReq := httptest.NewRequest(http.MethodGet, prefixProduct, nil)
		q := apiReq.URL.Query()
		q.Add("uid", uid.String())
		apiReq.URL.RawQuery = q.Encode()

		resp := &ds.GetProductResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}

		a.productMock.EXPECT().GetProduct(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetProduct(a.responseWriter, apiReq)
	})

	t.Run("GetProduct 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.GetProductRequest{
			Uid: uid,
		}

		apiReq := httptest.NewRequest(http.MethodGet, prefixProduct, nil)
		q := apiReq.URL.Query()
		q.Add("uid", uid.String())
		apiReq.URL.RawQuery = q.Encode()

		a.productMock.EXPECT().GetProduct(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetProduct(a.responseWriter, apiReq)
	})
}

func TestGetProducts(t *testing.T) {
	t.Parallel()

	t.Run("GetProducts 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		req := &ds.GetProductsRequest{
			Limit:  10,
			Offset: 1,
		}

		apiReq := httptest.NewRequest(http.MethodGet, prefixProducts, nil)
		q := apiReq.URL.Query()
		q.Add("limit", fmt.Sprint(req.Limit))
		q.Add("offset", fmt.Sprint(req.Offset))
		apiReq.URL.RawQuery = q.Encode()

		resp := &ds.GetProductsResponse{
			Products: []ds.Product{
				{
					Uid:             uuid.New(),
					SupplierUid:     uuid.New(),
					ImageUid:        uuid.New(),
					LastUpdateDate:  ds.DateOnlyFromString("01.01.2026"),
					Name:            "name",
					Category:        "category",
					Price:           299.99,
					AvaliableStocks: 20,
				},
			},
		}

		a.productMock.EXPECT().GetProducts(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetProducts(a.responseWriter, apiReq)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	t.Run("DeleteProduct 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodDelete, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		resp := &ds.DeleteProductResponse{}

		a.productMock.EXPECT().DeleteProduct(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DeleteProduct(a.responseWriter, apiReq)
	})

	t.Run("DeleteProduct 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		req := &ds.DeleteProductRequest{}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodDelete, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteProduct(a.responseWriter, apiReq)
	})

	t.Run("DeleteProduct 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodDelete, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		resp := &ds.DeleteProductResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}

		a.productMock.EXPECT().DeleteProduct(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DeleteProduct(a.responseWriter, apiReq)
	})

	t.Run("DeleteProduct 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		apiReq := httptest.NewRequest(http.MethodDelete, prefixProduct, strings.NewReader(string(jsonBody)))
		apiReq.Header.Set("Content-Type", "application/json")

		a.productMock.EXPECT().DeleteProduct(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteProduct(a.responseWriter, apiReq)
	})
}
