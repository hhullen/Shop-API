package postgres

import (
	"context"
	"database/sql"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddSupplier(t *testing.T) {
	t.Parallel()

	t.Run("AddSupplier ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddSupplierRequest{
			Supplier: ds.Supplier{
				Uid:         uid,
				Name:        "Name",
				PhoneNumber: "+79135468877 RU",
				Address: &ds.Address{
					Country: "USA",
					City:    "Redwood",
					Street:  "1st AVE",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().InsertSupplier(gomock.Any(), gomock.Any()).Return(uid, nil)

		resp, err := tc.client.AddSupplier(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("AddSupplier error on InsertAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddSupplierRequest{
			Supplier: ds.Supplier{
				Uid:         uid,
				Name:        "Name",
				PhoneNumber: "+79135468877 RU",
				Address: &ds.Address{
					Country: "USA",
					City:    "Redwood",
					Street:  "1st AVE",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.AddSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("AddSupplier error on InsertSupplier", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.AddSupplierRequest{
			Supplier: ds.Supplier{
				Uid:         uid,
				Name:        "Name",
				PhoneNumber: "+79135468877 RU",
				Address: &ds.Address{
					Country: "USA",
					City:    "Redwood",
					Street:  "1st AVE",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().InsertSupplier(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.AddSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestUpdateSupplierAddress(t *testing.T) {
	t.Parallel()

	t.Run("UpdateSupplierAddress ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uid, nil)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("UpdateSupplierAddress error on InsertAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("UpdateSupplierAddress error on UpdateSupplierAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("UpdateSupplierAddress not found", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(nil)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})

	t.Run("UpdateSupplierAddress not found error on CalculateSuppliersWithAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("UpdateSupplierAddress not found error on CalculateClientsWithAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("UpdateSupplierAddress not found error on DeleteAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.UpdateSupplierAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().UpdateSupplierAddress(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, sql.ErrNoRows)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(errTest)

		resp, err := tc.client.UpdateSupplierAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestDeleteSupplier(t *testing.T) {
	t.Parallel()

	t.Run("DeleteSupplier ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(nil)

		resp, err := tc.client.DeleteSupplier(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("DeleteSupplier not found", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(int32(0), sql.ErrNoRows)

		resp, err := tc.client.DeleteSupplier(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})

	t.Run("DeleteSupplier error on DeleteSupplier", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.DeleteSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteSupplier error on CalculateSuppliersWithAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DeleteSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteSupplier error on CalculateClientsWithAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DeleteSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteSupplier error on DeleteAddress", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.DeleteSupplierRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		addressId := int32(20)

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().DeleteSupplier(gomock.Any(), gomock.Any()).Return(addressId, nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(errTest)

		resp, err := tc.client.DeleteSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestGetSuppliers(t *testing.T) {
	t.Parallel()

	t.Run("GetSuppliers ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetSuppliersRequest{
			Limit:  10,
			Offset: 2,
		}

		res := []sqlc.SupplierDetail{
			{
				Uid:         uid,
				Name:        "name",
				PhoneNumber: "+79134568877 RU",
				Country:     "USA",
				City:        "Seattle",
				Street:      "Rob Ave",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetSuppliersPage(gomock.Any(), gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetSuppliers(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Suppliers[0].Uid, res[0].Uid)
		require.Equal(t, resp.Suppliers[0].Name, res[0].Name)
		require.Equal(t, resp.Suppliers[0].PhoneNumber, ds.PhoneNumber(res[0].PhoneNumber))
		require.Equal(t, resp.Suppliers[0].Address.Country, res[0].Country)
		require.Equal(t, resp.Suppliers[0].Address.City, res[0].City)
		require.Equal(t, resp.Suppliers[0].Address.Street, res[0].Street)
	})

	t.Run("GetSuppliers error on GetSuppliersPage", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		req := &ds.GetSuppliersRequest{
			Limit:  10,
			Offset: 2,
		}

		res := []sqlc.SupplierDetail{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetSuppliersPage(gomock.Any(), gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetSuppliers(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("GetSuppliers with NO offset and limit ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetSuppliersRequest{}

		res := []sqlc.SupplierDetail{
			{
				Uid:         uid,
				Name:        "name",
				PhoneNumber: "+79134568877 RU",
				Country:     "USA",
				City:        "Seattle",
				Street:      "Rob Ave",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllSuppliers(gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetSuppliers(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Suppliers[0].Uid, res[0].Uid)
		require.Equal(t, resp.Suppliers[0].Name, res[0].Name)
		require.Equal(t, resp.Suppliers[0].PhoneNumber, ds.PhoneNumber(res[0].PhoneNumber))
		require.Equal(t, resp.Suppliers[0].Address.Country, res[0].Country)
		require.Equal(t, resp.Suppliers[0].Address.City, res[0].City)
		require.Equal(t, resp.Suppliers[0].Address.Street, res[0].Street)
	})

	t.Run("GetSuppliers error on GetAllSuppliers", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		req := &ds.GetSuppliersRequest{}

		res := []sqlc.SupplierDetail{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllSuppliers(gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetSuppliers(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestGetSupplier(t *testing.T) {
	t.Parallel()

	t.Run("GetSupplier ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetSupplierRequest{
			Uid: uid,
		}

		res := sqlc.SupplierDetail{
			Uid:         uid,
			Name:        "name",
			PhoneNumber: "+79134568877 RU",
			Country:     "USA",
			City:        "Seattle",
			Street:      "Rob Ave",
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetSupplier(gomock.Any(), gomock.Any()).Return(res, nil)

		resp, err := tc.client.GetSupplier(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Supplier.Uid, res.Uid)
		require.Equal(t, resp.Supplier.Name, res.Name)
		require.Equal(t, resp.Supplier.PhoneNumber, ds.PhoneNumber(res.PhoneNumber))
		require.Equal(t, resp.Supplier.Address.Country, res.Country)
		require.Equal(t, resp.Supplier.Address.City, res.City)
		require.Equal(t, resp.Supplier.Address.Street, res.Street)
	})

	t.Run("GetSupplier error on GetSupplier", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.GetSupplierRequest{
			Uid: uid,
		}

		res := sqlc.SupplierDetail{}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetSupplier(gomock.Any(), gomock.Any()).Return(res, errTest)

		resp, err := tc.client.GetSupplier(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
