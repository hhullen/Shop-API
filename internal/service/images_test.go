package service

import (
	ds "shopapi/internal/datastruct"
	"testing"

	"github.com/golang/mock/gomock"
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

		s.imageStorageMock.EXPECT().GetProductImage(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetProductImage(req)
		require.NotNil(t, resp)
	})

	t.Run("GetProductImage error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetProductImageRequest{}

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

		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(res, nil)
		s.loggerMock.EXPECT().InfoKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.NotNil(t, resp)
	})

	t.Run("GetImage error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetImageRequest{}

		s.imageStorageMock.EXPECT().GetImage(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetImage(req)
		require.Nil(t, resp)
	})
}
