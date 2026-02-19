package service

import (
	ds "shopapi/internal/datastruct"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddSupplier(t *testing.T) {
	t.Parallel()

	t.Run("AddSupplier ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddSupplierRequest{}

		res := &ds.AddSupplierResponse{}

		s.supplierStorageMock.EXPECT().AddSupplier(gomock.Any()).Return(res, nil)

		resp := s.srv.AddSupplier(req)
		require.NotNil(t, resp)
	})

	t.Run("AddSupplier error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddSupplierRequest{}

		s.supplierStorageMock.EXPECT().AddSupplier(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.AddSupplier(req)
		require.Nil(t, resp)
	})
}

func TestUpdateSupplierAddress(t *testing.T) {
	t.Parallel()

	t.Run("UpdateSupplierAddress ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.UpdateSupplierAddressRequest{}

		res := &ds.UpdateSupplierAddressResponse{
			Status: ds.Status{Message: "status"},
		}

		s.supplierStorageMock.EXPECT().UpdateSupplierAddress(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.UpdateSupplierAddress(req)
		require.NotNil(t, resp)
	})

	t.Run("UpdateSupplierAddress error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.UpdateSupplierAddressRequest{}

		s.supplierStorageMock.EXPECT().UpdateSupplierAddress(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.UpdateSupplierAddress(req)
		require.Nil(t, resp)
	})
}

func TestDeleteSupplier(t *testing.T) {
	t.Parallel()

	t.Run("DeleteSupplier ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteSupplierRequest{}

		res := &ds.DeleteSupplierResponse{
			Status: ds.Status{Message: "status"},
		}

		s.supplierStorageMock.EXPECT().DeleteSupplier(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteSupplier(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteSupplier error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteSupplierRequest{}

		s.supplierStorageMock.EXPECT().DeleteSupplier(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteSupplier(req)
		require.Nil(t, resp)
	})
}

func TestGetSuppliers(t *testing.T) {
	t.Parallel()

	t.Run("GetSuppliers ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSuppliersRequest{}

		res := &ds.GetSuppliersResponse{}

		s.supplierStorageMock.EXPECT().GetSuppliers(gomock.Any()).Return(res, nil)

		resp := s.srv.GetSuppliers(req)
		require.NotNil(t, resp)
	})

	t.Run("GetSuppliers error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSuppliersRequest{}

		s.supplierStorageMock.EXPECT().GetSuppliers(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSuppliers(req)
		require.Nil(t, resp)
	})
}

func TestGetSupplier(t *testing.T) {
	t.Parallel()

	t.Run("GetSupplier ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{}

		res := &ds.GetSupplierResponse{
			Status: ds.Status{Message: "status"},
		}

		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
	})

	t.Run("GetSupplier error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{}

		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.Nil(t, resp)
	})
}
