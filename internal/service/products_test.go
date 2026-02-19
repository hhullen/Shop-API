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

		s.productStorageMock.EXPECT().AddProduct(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.AddProduct(req)
		require.NotNil(t, resp)
	})

	t.Run("AddProduct error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddProductRequest{}

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

		s.productStorageMock.EXPECT().GetProduct(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProduct(req)
		require.NotNil(t, resp)
	})

	t.Run("GetProduct error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductRequest{}

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

		s.productStorageMock.EXPECT().GetProducts(gomock.Any()).Return(res, nil)

		resp := s.srv.GetProducts(req)
		require.NotNil(t, resp)
	})

	t.Run("GetProducts error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductsRequest{}

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

		s.productStorageMock.EXPECT().DeleteProduct(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteProduct(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteProduct error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteProductRequest{}

		s.productStorageMock.EXPECT().DeleteProduct(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteProduct(req)
		require.Nil(t, resp)
	})
}
