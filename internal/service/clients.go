package service

import (
	ds "shopapi/internal/datastruct"
)

const (
	clientsCacheKey = "clients"
)

func (s *Service) AddClient(req *ds.AddClientRequest) *ds.AddClientResponse {
	resp, err := s.clientStorage.AddClient(req)
	if err != nil {
		s.logger.ErrorKV("failed on AddClient", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) DeleteClient(req *ds.DeleteClientRequest) *ds.DeleteClientResponse {
	resp, err := s.clientStorage.DeleteClient(req)
	if err != nil {
		s.logger.ErrorKV("failed on DeleteClient", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteClient", resp.GetStatus())

	return resp
}

func (s *Service) GetClients(req *ds.GetClientsRequest) *ds.GetClientsResponse {
	resp, err := s.clientStorage.GetClients(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetClients", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) GetClientsByName(req *ds.GetClientsByNameRequest) (resp *ds.GetClientsByNameResponse) {
	var cached bool
	var err error

	key := makeCacheKey(clientsCacheKey, req.Name, req.Surname)
	if !req.AvoidCache {
		resp = &ds.GetClientsByNameResponse{}
		cached, err = s.cache.Read(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed reading cache on GetClientsByName", "message", err.Error())
		}
	}

	if !cached {
		resp, err = s.clientStorage.GetClientsByName(req)
		if err != nil {
			s.logger.ErrorKV("failed on GetClientsByName", "message", err.Error())
			return nil
		}

		err = s.cache.Write(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed writing cache on GetClientsByName", "message", err.Error())
		}
		resp.Cached = false
	} else {
		resp.Cached = true
	}

	return
}

func (s *Service) PatchClientAddress(req *ds.PatchClientAddressRequest) *ds.PatchClientAddressResponse {
	resp, err := s.clientStorage.PatchClientAddress(req)
	if err != nil {
		s.logger.ErrorKV("failed on PatchClientAddress", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("PatchClientAddress", resp.GetStatus())

	return resp
}
