package service

import (
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddImage(t *testing.T) {
	t.Parallel()

	t.Run("AddImage ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddImageRequest{}

		res := &ds.AddImageResponse{}

		s.imageStorageMock.EXPECT().AddImage(gomock.Any()).Return(res, nil)

		resp := s.srv.AddImage(req)
		require.NotNil(t, resp)
	})

	t.Run("AddImage error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddImageRequest{}

		s.imageStorageMock.EXPECT().AddImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.AddImage(req)
		require.Nil(t, resp)
	})
}

func TestUpdateImage(t *testing.T) {
	t.Parallel()

	t.Run("UpdateImage ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.UpdateImageRequest{}

		res := &ds.UpdateImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.imageStorageMock.EXPECT().UpdateImage(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.UpdateImage(req)
		require.NotNil(t, resp)
	})

	t.Run("UpdateImage error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.UpdateImageRequest{}

		s.imageStorageMock.EXPECT().UpdateImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.UpdateImage(req)
		require.Nil(t, resp)
	})
}

func TestDeleteImage(t *testing.T) {
	t.Parallel()

	t.Run("DeleteImage ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteImageRequest{}

		res := &ds.DeleteImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.imageStorageMock.EXPECT().DeleteImage(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteImage(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteImage error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteImageRequest{}

		s.imageStorageMock.EXPECT().DeleteImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteImage(req)
		require.Nil(t, resp)
	})
}

func TestGetProductImage(t *testing.T) {
	t.Parallel()

	t.Run("GetProductImage ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{}

		res := &ds.GetProductImageResponse{
			Status: ds.Status{Message: "status"},
		}

		notCached := false

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProductImage avoid cache ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{
			AvoidCache: true,
		}

		res := &ds.GetProductImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProductImage cached ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{}

		uid := uuid.New()
		cached := true
		getCached := func(key string, v any) (bool, error) {
			vv := v.(*ds.GetProductImageResponse)
			vv.Image = supports.TestImage
			vv.Uid = &uid
			vv.Status = ds.Status{Message: "status"}
			return cached, nil
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(getCached)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
		require.True(t, resp.Cached)
		require.Equal(t, resp.Uid, &uid)
		require.Equal(t, resp.Image, supports.TestImage)
	})

	t.Run("GetProductImage error on Read", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{}

		res := &ds.GetProductImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProductImage error on Write", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{}

		res := &ds.GetProductImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetProductImage error on GetProductImage", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{AvoidCache: true}

		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.Nil(t, resp)
	})
}

func TestGetImage(t *testing.T) {
	t.Parallel()

	t.Run("GetImage ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{}

		res := &ds.GetImageResponse{
			Status: ds.Status{Message: "status"},
		}

		notCached := false

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetImage avoid cache ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{AvoidCache: true}

		res := &ds.GetImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetImage cached ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{}

		uid := uuid.New()
		cached := true
		getCached := func(key string, v any) (bool, error) {
			vv := v.(*ds.GetImageResponse)
			vv.Image = supports.TestImage
			vv.Uid = &uid
			vv.Status = ds.Status{Message: "status"}
			return cached, nil
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(getCached)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
		require.True(t, resp.Cached)
		require.Equal(t, resp.Uid, &uid)
		require.Equal(t, resp.Image, supports.TestImage)
	})

	t.Run("GetImage error on Read", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{}

		res := &ds.GetImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetImage error on Write", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{}

		res := &ds.GetImageResponse{
			Status: ds.Status{Message: "status"},
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, nil)
		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetImage error on GetImage", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{AvoidCache: true}

		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.Nil(t, resp)
	})
}
