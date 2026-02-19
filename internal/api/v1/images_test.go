package api

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPutImage(t *testing.T) {
	t.Parallel()

	t.Run("PutImage 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write(supports.TestImage)
		require.Nil(t, err)

		uid := uuid.New()
		err = writer.WriteField("uid", uid.String())
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPost, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		req := &ds.AddImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		resp := &ds.AddImageResponse{
			Uid: &uid,
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.imageMock.EXPECT().AddImage(req).Return(resp)

		a.api.PutImage(a.responseWriter, apiReq)
	})

	t.Run("PutImage 400 wrong image", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write([]byte("just s text file instead of image"))
		require.Nil(t, err)

		uid := uuid.New()
		err = writer.WriteField("uid", uid.String())
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPost, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.PutImage(a.responseWriter, apiReq)
	})

	t.Run("PutImage 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write(supports.TestImage)
		require.Nil(t, err)

		uid := uuid.New()
		err = writer.WriteField("uid", uid.String())
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPost, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		req := &ds.AddImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.imageMock.EXPECT().AddImage(req).Return(nil)

		a.api.PutImage(a.responseWriter, apiReq)
	})
}

func TestUpdateImage(t *testing.T) {
	t.Parallel()

	t.Run("UpdateImage 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write(supports.TestImage)
		require.Nil(t, err)

		uid := uuid.New()
		err = writer.WriteField("uid", uid.String())
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPatch, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		req := &ds.UpdateImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		resp := &ds.UpdateImageResponse{
			Status: ds.Status{Message: ds.StatusOK},
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			t.Fatal(err)
		}

		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(buf.Bytes())

		a.imageMock.EXPECT().UpdateImage(req).Return(resp)

		a.api.UpdateImage(a.responseWriter, apiReq)
	})

	t.Run("UpdateImage 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write(supports.TestImage)
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPatch, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.UpdateImage(a.responseWriter, apiReq)
	})

	t.Run("UpdateImage 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		buff := &bytes.Buffer{}

		writer := multipart.NewWriter(buff)

		part, _ := writer.CreateFormFile("image", "some.png")

		_, err := part.Write(supports.TestImage)
		require.Nil(t, err)

		uid := uuid.New()
		err = writer.WriteField("uid", uid.String())
		require.Nil(t, err)

		err = writer.Close()
		require.Nil(t, err)

		apiReq := httptest.NewRequest(http.MethodPatch, prefixImage, buff)

		apiReq.Header.Set("Content-Type", writer.FormDataContentType())

		req := &ds.UpdateImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.imageMock.EXPECT().UpdateImage(req).Return(nil)

		a.api.UpdateImage(a.responseWriter, apiReq)
	})
}

func TestGetProductImage(t *testing.T) {
	t.Parallel()

	t.Run("GetProductImage 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImageProduct, nil)

		productUid := uuid.New()

		q := testReq.URL.Query()
		q.Add("product_uid", productUid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetProductImageRequest{
			ProductUid: productUid,
		}

		imageUid := uuid.New()
		resp := &ds.GetProductImageResponse{
			Uid:   &imageUid,
			Image: supports.TestImage,
		}

		a.imageMock.EXPECT().GetProductImage(req).Return(resp)
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(supports.TestImage)

		a.api.GetProductImage(a.responseWriter, testReq)

	})

	t.Run("GetProductImage 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImageProduct, nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetProductImage(a.responseWriter, testReq)

	})

	t.Run("GetProductImage 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImageProduct, nil)

		productUid := uuid.New()

		q := testReq.URL.Query()
		q.Add("product_uid", productUid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetProductImageRequest{
			ProductUid: productUid,
		}

		resp := &ds.GetProductImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}

		a.imageMock.EXPECT().GetProductImage(req).Return(resp)
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetProductImage(a.responseWriter, testReq)

	})

	t.Run("GetProductImage 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImageProduct, nil)

		productUid := uuid.New()

		q := testReq.URL.Query()
		q.Add("product_uid", productUid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetProductImageRequest{
			ProductUid: productUid,
		}

		a.imageMock.EXPECT().GetProductImage(req).Return(nil)
		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetProductImage(a.responseWriter, testReq)

	})
}

func TestGetImage(t *testing.T) {
	t.Parallel()

	t.Run("GetImage 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImage, nil)

		uid := uuid.New()

		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetImageRequest{
			Uid: uid,
		}

		resp := &ds.GetImageResponse{
			Uid:   &uid,
			Image: supports.TestImage,
		}

		a.imageMock.EXPECT().GetImage(req).Return(resp)
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(supports.TestImage)

		a.api.GetImage(a.responseWriter, testReq)

	})

	t.Run("GetImage 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImage, nil)

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetImage(a.responseWriter, testReq)

	})

	t.Run("GetImage 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImage, nil)

		uid := uuid.New()

		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetImageRequest{
			Uid: uid,
		}

		resp := &ds.GetImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}

		a.imageMock.EXPECT().GetImage(req).Return(resp)
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetImage(a.responseWriter, testReq)

	})

	t.Run("GetImage 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		testReq := httptest.NewRequest(http.MethodGet, prefixImage, nil)

		uid := uuid.New()

		q := testReq.URL.Query()
		q.Add("uid", uid.String())
		testReq.URL.RawQuery = q.Encode()

		req := &ds.GetImageRequest{
			Uid: uid,
		}

		a.imageMock.EXPECT().GetImage(req).Return(nil)
		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.GetImage(a.responseWriter, testReq)

	})
}

func TestDeleteImage(t *testing.T) {
	t.Parallel()

	t.Run("DeleteImage 200", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteImageRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.imageMock.EXPECT().DeleteImage(reqStruct).Return(&ds.DeleteImageResponse{Status: ds.Status{Message: ds.StatusOK}})
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusOK)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteImage(a.responseWriter, testReq)
	})

	t.Run("DeleteImage 400", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteImageRequest{}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusBadRequest)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteImage(a.responseWriter, testReq)
	})

	t.Run("DeleteImage 404", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteImageRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.imageMock.EXPECT().DeleteImage(reqStruct).Return(&ds.DeleteImageResponse{Status: ds.Status{Message: ds.StatusNotFound}})
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusNotFound)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteImage(a.responseWriter, testReq)
	})

	t.Run("DeleteImage 500", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		a := NewTestApi(ctx, t)

		reqStruct := &ds.DeleteImageRequest{
			Uid: uuid.New(),
		}

		jsonBody, err := json.Marshal(&reqStruct)
		if err != nil {
			t.Fatal(err)
		}

		testReq := httptest.NewRequest(http.MethodDelete, prefixClient, strings.NewReader(string(jsonBody)))
		testReq.Header.Set("Content-Type", "application/json")

		a.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.Any())
		a.imageMock.EXPECT().DeleteImage(reqStruct).Return(nil)
		a.responseWriter.EXPECT().Header().Return(http.Header{}).MinTimes(1)
		a.responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
		a.responseWriter.EXPECT().Write(gomock.Any())

		a.api.DeleteImage(a.responseWriter, testReq)
	})
}
