package postgres

import (
	"context"
	"database/sql"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddProduct(t *testing.T) {
	t.Parallel()

	t.Run("AddProduct ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddProductRequest{
			Product: ds.Product{
				Uid:             uid,
				SupplierUid:     uuid.New(),
				ImageUid:        uuid.New(),
				LastUpdateDate:  ds.DateOnly(ds.DateOnlyFromString("10.12.2020")),
				Name:            "Name",
				Category:        "Category",
				Price:           299.95,
				AvaliableStocks: 123,
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().IsImageAndSupplierExists(gomock.Any(), gomock.Any()).Return(true, nil)
		tc.querierMock.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

		resp, err := tc.client.AddProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("AddProduct false on IsImageAndSupplierExists", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddProductRequest{
			Product: ds.Product{
				Uid:             uid,
				SupplierUid:     uuid.New(),
				ImageUid:        uuid.New(),
				LastUpdateDate:  ds.DateOnly(ds.DateOnlyFromString("10.12.2020")),
				Name:            "Name",
				Category:        "Category",
				Price:           299.95,
				AvaliableStocks: 123,
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().IsImageAndSupplierExists(gomock.Any(), gomock.Any()).Return(false, nil)

		resp, err := tc.client.AddProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusAddProductWithNoImageOrSupplier)
	})

	t.Run("AddProduct error on IsImageAndSupplierExists", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddProductRequest{
			Product: ds.Product{
				Uid:             uid,
				SupplierUid:     uuid.New(),
				ImageUid:        uuid.New(),
				LastUpdateDate:  ds.DateOnly(ds.DateOnlyFromString("10.12.2020")),
				Name:            "Name",
				Category:        "Category",
				Price:           299.95,
				AvaliableStocks: 123,
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().IsImageAndSupplierExists(gomock.Any(), gomock.Any()).Return(false, errTest)

		resp, err := tc.client.AddProduct(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("AddProduct error on InsertProduct", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddProductRequest{
			Product: ds.Product{
				Uid:             uid,
				SupplierUid:     uuid.New(),
				ImageUid:        uuid.New(),
				LastUpdateDate:  ds.DateOnly(ds.DateOnlyFromString("10.12.2020")),
				Name:            "Name",
				Category:        "Category",
				Price:           299.95,
				AvaliableStocks: 123,
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().IsImageAndSupplierExists(gomock.Any(), gomock.Any()).Return(true, nil)
		tc.querierMock.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.AddProduct(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestDecreaseProducts(t *testing.T) {
	t.Parallel()

	t.Run("DecreaseProducts ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: 10,
		}

		left := int64(20)
		shouldLeft := left - req.Amount

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().LockStockForUpdate(gomock.Any(), gomock.Any()).Return(left, nil)
		tc.querierMock.EXPECT().DecreaseProduct(gomock.Any(), gomock.Any()).Return(shouldLeft, nil)

		resp, err := tc.client.DecreaseProducts(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Left, &shouldLeft)
	})

	t.Run("DecreaseProducts not found", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: 10,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().LockStockForUpdate(gomock.Any(), gomock.Any()).Return(int64(0), sql.ErrNoRows)

		resp, err := tc.client.DecreaseProducts(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})

	t.Run("DecreaseProducts error on LockStockForUpdate", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: 10,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().LockStockForUpdate(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DecreaseProducts(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DecreaseProducts not enough", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: 30,
		}

		left := int64(20)
		shouldLeft := left

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().LockStockForUpdate(gomock.Any(), gomock.Any()).Return(left, nil)

		resp, err := tc.client.DecreaseProducts(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Left, &shouldLeft)
		require.Equal(t, resp.GetStatus(), ds.StatusDecreaseProductsFailed)
	})

	t.Run("DecreaseProducts error on DecreaseProduct", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DecreaseProductsRequest{
			Uid:    uid,
			Amount: 10,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().LockStockForUpdate(gomock.Any(), gomock.Any()).Return(int64(10), nil)
		tc.querierMock.EXPECT().DecreaseProduct(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DecreaseProducts(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	t.Run("GetProduct ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductRequest{
			Uid: uid,
		}

		updTime := time.Now()

		res := sqlc.Product{
			Uid:            uid,
			Name:           "Name",
			Category:       "category",
			Price:          29999,
			AvailableStock: 10,
			LastUpdateDate: updTime,
			SupplierID:     uuid.New(),
			ImageID:        uuid.New(),
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Product.Uid, res.Uid)
		require.Equal(t, resp.Product.Name, res.Name)
		require.Equal(t, resp.Product.Category, res.Category)
		require.Equal(t, resp.Product.Price, fromDBPrice(res.Price))
		require.Equal(t, resp.Product.AvaliableStocks, res.AvailableStock)
		require.Equal(t, resp.Product.LastUpdateDate, ds.DateOnly(res.LastUpdateDate))
		require.Equal(t, resp.Product.SupplierUid, res.SupplierID)
		require.Equal(t, resp.Product.ImageUid, res.ImageID)
	})

	t.Run("GetProduct not found", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductRequest{
			Uid: uid,
		}

		res := sqlc.Product{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Return(res, sql.ErrNoRows)

		resp, err := tc.client.GetProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)

	})

	t.Run("GetProduct error on GetProduct", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductRequest{
			Uid: uid,
		}

		res := sqlc.Product{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetProduct(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestGetProducts(t *testing.T) {
	t.Parallel()

	t.Run("GetProducts ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductsRequest{
			Limit:  10,
			Offset: 1,
		}

		updTime := time.Now()

		res := []sqlc.Product{
			{
				Uid:            uid,
				Name:           "Name",
				Category:       "category",
				Price:          29999,
				AvailableStock: 10,
				LastUpdateDate: updTime,
				SupplierID:     uuid.New(),
				ImageID:        uuid.New(),
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProductsPage(gomock.Any(), gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetProducts(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Products[0].Uid, res[0].Uid)
		require.Equal(t, resp.Products[0].Name, res[0].Name)
		require.Equal(t, resp.Products[0].Category, res[0].Category)
		require.Equal(t, resp.Products[0].Price, fromDBPrice(res[0].Price))
		require.Equal(t, resp.Products[0].AvaliableStocks, res[0].AvailableStock)
		require.Equal(t, resp.Products[0].LastUpdateDate, ds.DateOnly(res[0].LastUpdateDate))
		require.Equal(t, resp.Products[0].SupplierUid, res[0].SupplierID)
		require.Equal(t, resp.Products[0].ImageUid, res[0].ImageID)
	})

	t.Run("GetProducts error on GetProductsPage", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)
		req := &ds.GetProductsRequest{
			Limit:  10,
			Offset: 1,
		}

		res := []sqlc.Product{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetProductsPage(gomock.Any(), gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetProducts(req)
		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("GetProducts no offset and limit ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetProductsRequest{}

		updTime := time.Now()

		res := []sqlc.Product{
			{
				Uid:            uid,
				Name:           "Name",
				Category:       "category",
				Price:          29999,
				AvailableStock: 10,
				LastUpdateDate: updTime,
				SupplierID:     uuid.New(),
				ImageID:        uuid.New(),
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllProducts(gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetProducts(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Products[0].Uid, res[0].Uid)
		require.Equal(t, resp.Products[0].Name, res[0].Name)
		require.Equal(t, resp.Products[0].Category, res[0].Category)
		require.Equal(t, resp.Products[0].Price, fromDBPrice(res[0].Price))
		require.Equal(t, resp.Products[0].AvaliableStocks, res[0].AvailableStock)
		require.Equal(t, resp.Products[0].LastUpdateDate, ds.DateOnly(res[0].LastUpdateDate))
		require.Equal(t, resp.Products[0].SupplierUid, res[0].SupplierID)
		require.Equal(t, resp.Products[0].ImageUid, res[0].ImageID)
	})

	t.Run("GetProducts no offset and limit error on GetAllProducts", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		req := &ds.GetProductsRequest{}

		res := []sqlc.Product{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllProducts(gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetProducts(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	t.Run("DeleteProduct ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(uid, nil)

		resp, err := tc.client.DeleteProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)

	})

	t.Run("DeleteProduct not found", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)

		resp, err := tc.client.DeleteProduct(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)

	})

	t.Run("DeleteProduct error on DeleteProduct", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteProductRequest{
			Uid: uid,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.DeleteProduct(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
