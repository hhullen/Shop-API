package service

import (
	ds "shopapi/internal/datastruct"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddProduct(t *testing.T) {
	t.Parallel()

	t.Run("AddProduct ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddProductRequest{}

		res := &ds.AddProductResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.productStorageMock.EXPECT().AddProduct(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.AddProduct(req)
		require.NotNil(t, resp)
	})

	t.Run("AddProduct error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddProductRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.productStorageMock.EXPECT().AddProduct(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.AddProduct(req)
		require.Nil(t, resp)
	})
}

func TestDecreaseProducts(t *testing.T) {
	t.Parallel()

	t.Run("DecreaseProducts ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DecreaseProductsRequest{}

		res := &ds.DecreaseProductsResponse{
			Status: ds.Status{Message: "status"},
		}

		s.productStorageMock.EXPECT().DecreaseProducts(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DecreaseProducts(req)
		require.NotNil(t, resp)
	})

	t.Run("DecreaseProducts error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DecreaseProductsRequest{}

		s.productStorageMock.EXPECT().DecreaseProducts(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DecreaseProducts(req)
		require.Nil(t, resp)
	})
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	t.Run("GetProduct ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{}

		res := &ds.GetProductResponse{
			Status: ds.Status{Message: "status"},
		}

		notCached := false

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProduct cached ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{}

		cached := true

		name := "name"
		category := "category"
		getCached := func(key string, v any) (bool, error) {
			target := v.(**ds.GetProductResponse)
			vv := &ds.GetProductResponse{}
			vv.Product = &ds.Product{
				Name:     name,
				Category: category,
			}
			vv.Status = ds.Status{Message: "status"}

			*target = vv
			return cached, nil
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(getCached)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
		require.True(t, resp.Cached)
		require.Equal(t, resp.Product.Name, name)
		require.Equal(t, resp.Product.Category, category)
	})

	t.Run("GetProduct avoid cache ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{AvoidCacheFlag: ds.AvoidCacheFlag{Flag: true}}

		res := &ds.GetProductResponse{
			Status: ds.Status{Message: "status"},
		}

		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProduct error on Read", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{}

		res := &ds.GetProductResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProduct error on Write", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{}

		res := &ds.GetProductResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProduct error on GetProduct", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{AvoidCacheFlag: ds.AvoidCacheFlag{Flag: true}}

		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.Nil(t, resp)
	})
}

func TestGetProducts(t *testing.T) {
	t.Parallel()

	t.Run("GetProducts ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductsRequest{}

		res := &ds.GetProductsResponse{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.productStorageMock.EXPECT().GetProducts(gomock.Any()).Return(res, nil)

		resp := s.srv.GetProducts(req)
		require.NotNil(t, resp)
	})

	t.Run("GetProducts error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductsRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.productStorageMock.EXPECT().GetProducts(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProducts(req)
		require.Nil(t, resp)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	t.Run("DeleteProduct ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteProductRequest{}

		res := &ds.DeleteProductResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.productStorageMock.EXPECT().DeleteProduct(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteProduct(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteProduct error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteProductRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.productStorageMock.EXPECT().DeleteProduct(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteProduct(req)
		require.Nil(t, resp)
	})
}
