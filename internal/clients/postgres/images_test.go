package postgres

import (
	"context"
	"database/sql"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddImage(t *testing.T) {
	t.Parallel()

	t.Run("AddImage Ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().AddImage(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

		resp, err := tc.client.AddImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("AddImage error on AddImage", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().AddImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.AddImage(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("AddImage already exsists", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().AddImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)

		resp, err := tc.client.AddImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusAlreadyExists)

	})
}

func TestUpdateImage(t *testing.T) {
	t.Parallel()

	t.Run("UpdateImage Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().UpdateImage(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

		resp, err := tc.client.UpdateImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("UpdateImage error on UpdateImage", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().UpdateImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.UpdateImage(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("UpdateImage not found", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateImageRequest{
			Uid:   uid,
			Image: supports.TestImage,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().UpdateImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)

		resp, err := tc.client.UpdateImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})
}

func TestDeleteImage(t *testing.T) {
	t.Parallel()

	t.Run("DeleteImage Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteImage(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

		resp, err := tc.client.DeleteImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("DeleteImage error on DeleteImage", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.DeleteImage(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteImage not found", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteImage(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)

		resp, err := tc.client.DeleteImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})
}

func TestGetProductImage(t *testing.T) {
	t.Parallel()

	t.Run("GetProductImage Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductImageRequest{
			ProductUid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProductImage(gomock.Any(), gomock.Any()).Return(sqlc.Image{
			Uid:   uuid.New(),
			Image: supports.TestImage,
		}, nil)

		resp, err := tc.client.GetProductImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("GetProductImage error on GetProductImage", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductImageRequest{
			ProductUid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProductImage(gomock.Any(), gomock.Any()).Return(sqlc.Image{}, errTest)

		resp, err := tc.client.GetProductImage(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("GetProductImage not found", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductImageRequest{
			ProductUid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProductImage(gomock.Any(), gomock.Any()).Return(sqlc.Image{}, sql.ErrNoRows)

		resp, err := tc.client.GetProductImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})
}

func TestGetImage(t *testing.T) {
	t.Parallel()

	t.Run("GetImage Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetImage(gomock.Any(), gomock.Any()).Return(supports.TestImage, nil)

		resp, err := tc.client.GetImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("GetImage error on GetImage", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetImage(gomock.Any(), gomock.Any()).Return(nil, errTest)

		resp, err := tc.client.GetImage(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("GetImage not found", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetImageRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetImage(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)

		resp, err := tc.client.GetImage(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})
}
