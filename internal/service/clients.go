package service

import (
	"encoding/json"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"strconv"
)

func (s *Service) AddClient(req *ds.AddClientRequest) *ds.AddClientResponse {
	reqKey, err := json.Marshal(req.Client)
	if err != nil {
		s.logger.ErrorKV("failed making cache key", "message", err.Error(), "data", *req)
		req.AvoidCacheFlag.Flag = true
	}
	key := makeCacheKey("AddClient", supports.GetHash(reqKey))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.AddClientResponse, error) {
		return s.clientStorage.AddClient(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on AddClient", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) DeleteClient(req *ds.DeleteClientRequest) *ds.DeleteClientResponse {
	key := makeCacheKey("DeleteClient", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.DeleteClientResponse, error) {
		return s.clientStorage.DeleteClient(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on DeleteClient", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteClient", resp.GetStatus())

	return resp
}

func (s *Service) GetClients(req *ds.GetClientsRequest) *ds.GetClientsResponse {
	key := makeCacheKey("GetClients", strconv.FormatInt(req.Limit, 10), strconv.FormatInt(req.Offset, 10))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetClientsResponse, error) {
		return s.clientStorage.GetClients(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on GetClients", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) GetClientsByName(req *ds.GetClientsByNameRequest) *ds.GetClientsByNameResponse {
	key := makeCacheKey("GetClientsByName", req.Name, req.Surname)

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetClientsByNameResponse, error) {
		return s.clientStorage.GetClientsByName(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on GetClientsByName", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) PatchClientAddress(req *ds.PatchClientAddressRequest) *ds.PatchClientAddressResponse {
	reqKey, err := json.Marshal(req.Address)
	if err != nil {
		s.logger.ErrorKV("failed making cache key", "message", err.Error(), "data", *req)
		req.AvoidCacheFlag.Flag = true
	}
	key := makeCacheKey("PatchClientAddress", req.Uid.String(), supports.GetHash(reqKey))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.PatchClientAddressResponse, error) {
		return s.clientStorage.PatchClientAddress(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on PatchClientAddress", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("PatchClientAddress", resp.GetStatus())

	return resp
}
