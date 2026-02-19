package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddClient(t *testing.T) {
	t.Parallel()

	t.Run("AddClient Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.AddClientRequest{
			Client: ds.Client{
				Uid:              uid,
				Birthday:         ds.DateOnlyFromString("10.12.2011"),
				RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
				Name:             "Vasilisa",
				Surname:          "Kadyk",
				Gender:           ds.Male,
				Address: &ds.Address{
					Country: "USA",
					City:    "Seattle",
					Street:  "12th Ave E",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(int32(1), nil)
		tc.querierMock.EXPECT().InsertClient(gomock.Any(), gomock.Any()).Return(uid, nil)

		resp, err := tc.client.AddClient(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("AddClient error on InsertAddress", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.AddClientRequest{
			Client: ds.Client{
				Uid:              uid,
				Birthday:         ds.DateOnlyFromString("10.12.2011"),
				RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
				Name:             "Vasilisa",
				Surname:          "Kadyk",
				Gender:           ds.Male,
				Address: &ds.Address{
					Country: "USA",
					City:    "Seattle",
					Street:  "12th Ave E",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.AddClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("AddClient error on InsertClient", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.AddClientRequest{
			Client: ds.Client{
				Uid:              uid,
				Birthday:         ds.DateOnlyFromString("10.12.2011"),
				RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
				Name:             "Vasilisa",
				Surname:          "Kadyk",
				Gender:           ds.Male,
				Address: &ds.Address{
					Country: "USA",
					City:    "Seattle",
					Street:  "12th Ave E",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)
		tc.querierMock.EXPECT().InsertAddress(gomock.Any(), gomock.Any()).Return(int32(1), nil)
		tc.querierMock.EXPECT().InsertClient(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errTest)

		resp, err := tc.client.AddClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("AddClient error on ExecTx", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.AddClientRequest{
			Client: ds.Client{
				Uid:              uid,
				Birthday:         ds.DateOnlyFromString("10.12.2011"),
				RegistrationDate: ds.DateOnlyFromString("30/01/2026"),
				Name:             "Vasilisa",
				Surname:          "Kadyk",
				Gender:           ds.Male,
				Address: &ds.Address{
					Country: "USA",
					City:    "Seattle",
					Street:  "12th Ave E",
				},
			},
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return errTest
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		resp, err := tc.client.AddClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestDeleteClient(t *testing.T) {
	t.Parallel()

	t.Run("DeleteClient Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(11), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(nil)

		resp, err := tc.client.DeleteClient(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("DeleteClient error on DeleteClient", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.DeleteClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteClient error on CalculateClientsWithAddress", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(11), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DeleteClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteClient error on CalculateSuppliersWithAddress", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(11), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), errTest)

		resp, err := tc.client.DeleteClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteClient error on DeleteAddress", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(11), nil)
		tc.querierMock.EXPECT().CalculateClientsWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().CalculateSuppliersWithAddress(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		tc.querierMock.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(errTest)

		resp, err := tc.client.DeleteClient(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("DeleteClient not found", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()

		req := &ds.DeleteClientRequest{
			Uid: uid,
		}

		txExec := func(opts *sql.TxOptions, fn func(ctx context.Context, q IQuerier) error) error {
			return fn(tc.ctx, tc.querierMock)
		}

		tc.clientMock.EXPECT().ExecTx(gomock.Any(), gomock.Any()).DoAndReturn(txExec)

		tc.querierMock.EXPECT().DeleteClient(gomock.Any(), gomock.Any()).Return(int32(0), sql.ErrNoRows)

		resp, err := tc.client.DeleteClient(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.GetStatus(), ds.StatusNotFound)
	})
}

func TestGetClients(t *testing.T) {
	t.Parallel()

	t.Run("GetClients with NO offset and limit Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsRequest{
			Limit:  0,
			Offset: 0,
		}
		uid := uuid.New()

		sqlcResp := []sqlc.ClientDetail{
			{
				ClientName:       "name",
				ClientSurname:    "surname",
				Birthday:         time.Time(ds.DateOnlyFromString("13.08.1995")),
				Gender:           string(ds.Male),
				Uid:              uid,
				RegistrationDate: time.Time(ds.DateOnlyFromString("13.08.1995")),
				Country:          "USA",
				City:             "Redwood",
				Street:           "1st AVE",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllClients(gomock.Any()).Return(sqlcResp, nil)

		resp, err := tc.client.GetClients(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Clients[0].Name, sqlcResp[0].ClientName)
		require.Equal(t, resp.Clients[0].Surname, sqlcResp[0].ClientSurname)
		require.Equal(t, resp.Clients[0].Birthday, ds.DateOnly(sqlcResp[0].Birthday))
		require.Equal(t, resp.Clients[0].Gender, ds.Gender(sqlcResp[0].Gender))
		require.Equal(t, resp.Clients[0].Uid, sqlcResp[0].Uid)
		require.Equal(t, resp.Clients[0].RegistrationDate, ds.DateOnly(sqlcResp[0].RegistrationDate))
		require.Equal(t, resp.Clients[0].Address.Country, sqlcResp[0].Country)
		require.Equal(t, resp.Clients[0].Address.City, sqlcResp[0].City)
		require.Equal(t, resp.Clients[0].Address.Street, sqlcResp[0].Street)
	})

	t.Run("GetClients with NO offset and limit error", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsRequest{
			Limit:  0,
			Offset: 0,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetAllClients(gomock.Any()).Return(nil, errTest)

		resp, err := tc.client.GetClients(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("GetClients with offset and limit Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsRequest{
			Limit:  4,
			Offset: 1,
		}

		uid := uuid.New()
		sqlcResp := []sqlc.ClientDetail{
			{
				ClientName:       "name",
				ClientSurname:    "surname",
				Birthday:         time.Time(ds.DateOnlyFromString("13.08.1995")),
				Gender:           string(ds.Male),
				Uid:              uid,
				RegistrationDate: time.Time(ds.DateOnlyFromString("13.08.1995")),
				Country:          "USA",
				City:             "Redwood",
				Street:           "1st AVE",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetClientsPage(gomock.Any(), sqlc.GetClientsPageParams{
			Offset: int32(req.Offset),
			Limit:  int32(req.Limit),
		}).Return(sqlcResp, nil)

		resp, err := tc.client.GetClients(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Clients[0].Name, sqlcResp[0].ClientName)
		require.Equal(t, resp.Clients[0].Surname, sqlcResp[0].ClientSurname)
		require.Equal(t, resp.Clients[0].Birthday, ds.DateOnly(sqlcResp[0].Birthday))
		require.Equal(t, resp.Clients[0].Gender, ds.Gender(sqlcResp[0].Gender))
		require.Equal(t, resp.Clients[0].Uid, sqlcResp[0].Uid)
		require.Equal(t, resp.Clients[0].RegistrationDate, ds.DateOnly(sqlcResp[0].RegistrationDate))
		require.Equal(t, resp.Clients[0].Address.Country, sqlcResp[0].Country)
		require.Equal(t, resp.Clients[0].Address.City, sqlcResp[0].City)
		require.Equal(t, resp.Clients[0].Address.Street, sqlcResp[0].Street)
	})

	t.Run("GetClients with offset and limit Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsRequest{
			Limit:  4,
			Offset: 1,
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetClientsPage(gomock.Any(), sqlc.GetClientsPageParams{
			Offset: int32(req.Offset),
			Limit:  int32(req.Limit),
		}).Return(nil, errTest)

		resp, err := tc.client.GetClients(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestGetClientsByName(t *testing.T) {
	t.Parallel()

	t.Run("GetClientsByName Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsByNameRequest{
			Name:    "Name",
			Surname: "Surname",
		}
		uid := uuid.New()

		sqlcResp := []sqlc.ClientDetail{
			{
				ClientName:       req.Name,
				ClientSurname:    req.Surname,
				Birthday:         time.Time(ds.DateOnlyFromString("13.08.1995")),
				Gender:           string(ds.Male),
				Uid:              uid,
				RegistrationDate: time.Time(ds.DateOnlyFromString("13.08.1995")),
				Country:          "USA",
				City:             "Redwood",
				Street:           "1st AVE",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetClientsWithName(gomock.Any(), gomock.Any()).Return(sqlcResp, nil)

		resp, err := tc.client.GetClientsByName(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.Clients[0].Name, sqlcResp[0].ClientName)
		require.Equal(t, resp.Clients[0].Surname, sqlcResp[0].ClientSurname)
	})

	t.Run("GetClientsByName error on GetClientsWithName", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		req := &ds.GetClientsByNameRequest{
			Name:    "Name",
			Surname: "Surname",
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().GetClientsWithName(gomock.Any(), gomock.Any()).Return(nil, errTest)

		resp, err := tc.client.GetClientsByName(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}

func TestPatchClientAddress(t *testing.T) {
	t.Parallel()

	t.Run("PatchClientAddress Ok", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.PatchClientAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().UpdateClientAddress(gomock.Any(), gomock.Any()).Return(int32(0), nil)

		resp, err := tc.client.PatchClientAddress(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("PatchClientAddress error on UpdateClientAddress", func(t *testing.T) {
		t.Parallel()
		tc := NewTestClient(t)

		uid := uuid.New()
		req := &ds.PatchClientAddressRequest{
			Uid: uid,
			Address: &ds.Address{
				Country: "USA",
				City:    "Redwood",
				Street:  "1st AVE",
			},
		}

		tc.clientMock.EXPECT().CtxWithCancel().Return(context.Background(), func() {})
		tc.clientMock.EXPECT().Querier().Return(tc.querierMock)
		tc.querierMock.EXPECT().UpdateClientAddress(gomock.Any(), gomock.Any()).Return(int32(0), errTest)

		resp, err := tc.client.PatchClientAddress(req)
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
