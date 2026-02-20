package service

import (
	ds "shopapi/internal/datastruct"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddClient(t *testing.T) {
	t.Parallel()

	t.Run("AddClient ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddClientRequest{}

		res := &ds.AddClientResponse{}

		s.clientStorageMock.EXPECT().AddClient(gomock.Any()).Return(res, nil)

		resp := s.srv.AddClient(req)
		require.NotNil(t, resp)
	})

	t.Run("AddClient error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.AddClientRequest{}

		s.clientStorageMock.EXPECT().AddClient(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.AddClient(req)
		require.Nil(t, resp)
	})
}

func TestDeleteClient(t *testing.T) {
	t.Parallel()

	t.Run("DeleteClient ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteClientRequest{}

		res := &ds.DeleteClientResponse{}

		s.clientStorageMock.EXPECT().DeleteClient(gomock.Any()).Return(res, nil)

		resp := s.srv.DeleteClient(req)
		require.NotNil(t, resp)
	})

	t.Run("DeleteClient error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.DeleteClientRequest{}

		s.clientStorageMock.EXPECT().DeleteClient(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.DeleteClient(req)
		require.Nil(t, resp)
	})
}

func TestGetClients(t *testing.T) {
	t.Parallel()

	t.Run("GetClients ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsRequest{}

		res := &ds.GetClientsResponse{}

		s.clientStorageMock.EXPECT().GetClients(gomock.Any()).Return(res, nil)

		resp := s.srv.GetClients(req)
		require.NotNil(t, resp)
	})

	t.Run("GetClients error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsRequest{}

		s.clientStorageMock.EXPECT().GetClients(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetClients(req)
		require.Nil(t, resp)
	})
}

func TestGetClientsByName(t *testing.T) {
	t.Parallel()

	t.Run("GetClientsByName ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{}

		res := &ds.GetClientsByNameResponse{}

		notCached := false
		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.clientStorageMock.EXPECT().GetClientsByName(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)

		resp := s.srv.GetClientsByName(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetClientsByName cached ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{}

		res := &ds.GetClientsByNameResponse{}

		name := "Name"
		surname := "Surname"
		cached := true
		getCached := func(key string, v any) (bool, error) {
			vv := v.(*ds.GetClientsByNameResponse)
			vv.Clients = append(vv.Clients, ds.Client{
				Name:    name,
				Surname: surname,
			})
			return cached, nil
		}

		s.cacheMock.EXPECT().Read(gomock.Any(), res).DoAndReturn(getCached)

		resp := s.srv.GetClientsByName(req)
		require.NotNil(t, resp)
		require.True(t, resp.Cached)
		require.Equal(t, resp.Clients[0].Name, name)
		require.Equal(t, resp.Clients[0].Surname, surname)
	})

	t.Run("GetClientsByName avoid cache ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{
			AvoidCache: true,
		}

		res := &ds.GetClientsByNameResponse{}

		s.clientStorageMock.EXPECT().GetClientsByName(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)

		resp := s.srv.GetClientsByName(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached)
	})

	t.Run("GetClientsByName error on Read", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{}

		res := &ds.GetClientsByNameResponse{}

		notCached := false
		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(false, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())
		s.clientStorageMock.EXPECT().GetClientsByName(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)

		resp := s.srv.GetClientsByName(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached, notCached)
	})

	t.Run("GetClientsByName error on Write", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{}

		res := &ds.GetClientsByNameResponse{}

		notCached := false
		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.clientStorageMock.EXPECT().GetClientsByName(gomock.Any()).Return(res, nil)
		s.cacheMock.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetClientsByName(req)
		require.NotNil(t, resp)
		require.False(t, resp.Cached, notCached)
	})

	t.Run("GetClientsByName error on GetClientsByName", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.GetClientsByNameRequest{}

		notCached := false
		s.cacheMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(notCached, nil)
		s.clientStorageMock.EXPECT().GetClientsByName(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.GetClientsByName(req)
		require.Nil(t, resp)
	})
}

func TestPatchClientAddress(t *testing.T) {
	t.Parallel()

	t.Run("PatchClientAddress ok", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.PatchClientAddressRequest{}

		res := &ds.PatchClientAddressResponse{}

		s.clientStorageMock.EXPECT().PatchClientAddress(gomock.Any()).Return(res, nil)

		resp := s.srv.PatchClientAddress(req)
		require.NotNil(t, resp)
	})

	t.Run("PatchClientAddress error", func(t *testing.T) {
		t.Parallel()

		s := NewTestService(t)

		req := &ds.PatchClientAddressRequest{}

		s.clientStorageMock.EXPECT().PatchClientAddress(gomock.Any()).Return(nil, errTest)
		s.loggerMock.EXPECT().ErrorKV(gomock.Any(), gomock.All())

		resp := s.srv.PatchClientAddress(req)
		require.Nil(t, resp)
	})
}
