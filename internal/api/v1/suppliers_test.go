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

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestPutSupplier(t *testing.T) {
	t.Parallel()

	t.Run("PutSupplier 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Supplier{
			Uid:         uid,
			PhoneNumber: "+79994567722 RU",
			Name:        "Vasilisa&co",
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

		testReq := httptest.NewRequest(http.MethodPut, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.AddSupplierResponse{
			Uid: &uid,
		}

		a.supplierMock.EXPECT().AddSupplier(&ds.AddSupplierRequest{
			Supplier: clientStruct,
		}).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.PutSupplier(a.responseWriter, testReq)
	})

	t.Run("PutSupplier 400 no field", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Supplier{
			Uid:  uid,
			Name: "Vasilisa&co",
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

		testReq := httptest.NewRequest(http.MethodPut, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutSupplier(a.responseWriter, testReq)
	})

	t.Run("PutSupplier 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()

		clientStruct := ds.Supplier{
			Uid:         uid,
			PhoneNumber: "+79994567722 RU",
			Name:        "Vasilisa&co",
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

		testReq := httptest.NewRequest(http.MethodPut, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.supplierMock.EXPECT().AddSupplier(&ds.AddSupplierRequest{
			Supplier: clientStruct,
		}).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutSupplier(a.responseWriter, testReq)
	})
}

func TestUpdateSupplierAddress(t *testing.T) {
	t.Parallel()

	t.Run("UpdateSupplierAddress 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixSupplierAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.UpdateSupplierAddressResponse{Status: ds.Status{Message: ds.StatusOK}}

		a.supplierMock.EXPECT().UpdateSupplierAddress(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.UpdateSupplierAddress(a.responseWriter, testReq)
	})

	t.Run("UpdateSupplierAddress 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixSupplierAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.UpdateSupplierAddress(a.responseWriter, testReq)
	})

	t.Run("UpdateSupplierAddress 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixSupplierAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.UpdateSupplierAddressResponse{Status: ds.Status{Message: ds.StatusNotFound}}

		a.supplierMock.EXPECT().UpdateSupplierAddress(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.UpdateSupplierAddress(a.responseWriter, testReq)
	})

	t.Run("UpdateSupplierAddress 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodPatch, prefixSupplierAddress, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.supplierMock.EXPECT().UpdateSupplierAddress(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.UpdateSupplierAddress(a.responseWriter, testReq)
	})
}

func TestDeleteSupplier(t *testing.T) {
	t.Parallel()

	t.Run("DeleteSupplier 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.DeleteSupplierResponse{Status: ds.Status{Message: ds.StatusOK}}

		a.supplierMock.EXPECT().DeleteSupplier(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DeleteSupplier(a.responseWriter, testReq)
	})

	t.Run("DeleteSupplier 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		req := &ds.DeleteSupplierRequest{}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteSupplier(a.responseWriter, testReq)
	})

	t.Run("DeleteSupplier 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		resp := &ds.DeleteSupplierResponse{Status: ds.Status{Message: ds.StatusNotFound}}

		a.supplierMock.EXPECT().DeleteSupplier(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.DeleteSupplier(a.responseWriter, testReq)
	})

	t.Run("DeleteSupplier 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		jsonBody, err := json.Marshal(&req)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixSupplier, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.supplierMock.EXPECT().DeleteSupplier(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteSupplier(a.responseWriter, testReq)
	})
}

func TestGetSupplier(t *testing.T) {
	t.Parallel()

	t.Run("GetSupplier 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.GetSupplierRequest{
			Uid: uid,
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixSupplier, nil)
		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetSupplierResponse{Supplier: &ds.Supplier{
			Uid:         uid,
			PhoneNumber: "+79994567722 RU",
			Name:        "Vasilisa&co",
			Address: &ds.Address{
				Country: "USA",
				City:    "Seattle",
				Street:  "12th Ave E",
			},
		}}

		a.supplierMock.EXPECT().GetSupplier(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetSupplier(a.responseWriter, testReq)
	})

	t.Run("GetSupplier 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixSupplier, nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetSupplier(a.responseWriter, testReq)
	})

	t.Run("GetSupplier 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.GetSupplierRequest{
			Uid: uid,
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixSupplier, nil)
		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetSupplierResponse{Status: ds.Status{Message: ds.StatusNotFound}}

		a.supplierMock.EXPECT().GetSupplier(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetSupplier(a.responseWriter, testReq)
	})

	t.Run("GetSupplier 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.GetSupplierRequest{
			Uid: uid,
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixSupplier, nil)
		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		a.supplierMock.EXPECT().GetSupplier(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetSupplier(a.responseWriter, testReq)
	})
}

func TestGetSuppliers(t *testing.T) {
	t.Parallel()

	t.Run("GetSuppliers 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		uid := uuid.New()
		req := &ds.GetSuppliersRequest{
			Limit:  10,
			Offset: 1,
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixSuppliers, nil)
		q := testReq.URL.Query()
		q.Add("limit", fmt.Sprint(req.Limit))
		q.Add("offset", fmt.Sprint(req.Offset))
		testReq.URL.RawQuery = q.Encode()

		resp := &ds.GetSuppliersResponse{Suppliers: []ds.Supplier{
			{
				Uid:         uid,
				PhoneNumber: "+79994567722 RU",
				Name:        "Vasilisa&co",
				Address: &ds.Address{
					Country: "USA",
					City:    "Seattle",
					Street:  "12th Ave E",
				},
			},
		}}

		a.supplierMock.EXPECT().GetSuppliers(req).Return(resp)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.api.GetSuppliers(a.responseWriter, testReq)
	})

	t.Run("GetSuppliers 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		req := &ds.GetSuppliersRequest{
			Limit:  10,
			Offset: 1,
		}

		testReq := httptest.NewRequest(http.MethodGet, prefixSuppliers, nil)
		q := testReq.URL.Query()
		q.Add("limit", fmt.Sprint(req.Limit))
		q.Add("offset", fmt.Sprint(req.Offset))
		testReq.URL.RawQuery = q.Encode()

		a.supplierMock.EXPECT().GetSuppliers(req).Return(nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetSuppliers(a.responseWriter, testReq)
	})
}
