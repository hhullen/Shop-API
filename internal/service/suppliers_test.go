package service

import (
	ds "shopapi/internal/datastruct"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddSupplier(t *testing.T) {
	t.Parallel()

	t.Run("AddSupplier ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddSupplierRequest{}

		res := &ds.AddSupplierResponse{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.supplierStorageMock.EXPECT().AddSupplier(gomock.Any()).Return(res, nil)

		resp := s.srv.AddSupplier(req)
		require.NotNil(t, resp)
	})

	t.Run("AddSupplier error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddSupplierRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
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

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.supplierStorageMock.EXPECT().UpdateSupplierAddress(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.UpdateSupplierAddress(req)
		require.NotNil(t, resp)
	})

	t.Run("UpdateSupplierAddress error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.UpdateSupplierAddressRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
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

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.supplierStorageMock.EXPECT().DeleteSupplier(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteSupplier(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteSupplier error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteSupplierRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
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

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.supplierStorageMock.EXPECT().GetSuppliers(gomock.Any()).Return(res, nil)

		resp := s.srv.GetSuppliers(req)
		require.NotNil(t, resp)
	})

	t.Run("GetSuppliers error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSuppliersRequest{}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
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

		notCached := false

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetSupplier avoid cache ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{AvoidCacheFlag: ds.AvoidCacheFlag{Flag: true}}

		res := &ds.GetSupplierResponse{
			Status: ds.Status{Message: "status"},
		}

		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetSupplier cached ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{}

		cached := true
		name := "name"
		uid := uuid.New()
		getCached := func(key string, v any) (bool, error) {
			target := v.(**ds.GetSupplierResponse)
			vv := &ds.GetSupplierResponse{}
			vv.Supplier = &ds.Supplier{
				Name: name,
				Uid:  uid,
			}
			vv.Status = ds.Status{Message: "status"}
			*target = vv
			return cached, nil
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(getCached)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
		require.True(t, resp.Cached)
		require.Equal(t, resp.Supplier.Name, name)
		require.Equal(t, resp.Supplier.Uid, uid)
	})

	t.Run("GetSupplier error on Read", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{}

		res := &ds.GetSupplierResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetSupplier error on Write", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{}

		res := &ds.GetSupplierResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetSupplier error on GetSupplier", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetSupplierRequest{AvoidCacheFlag: ds.AvoidCacheFlag{Flag: true}}

		s.supplierStorageMock.EXPECT().GetSupplier(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetSupplier(req)
		require.Nil(t, resp)
	})
}
