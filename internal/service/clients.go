package service

import (
	ds "shopapi/internal/datastruct"
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

func (s *Service) GetClientsByName(req *ds.GetClientsByNameRequest) *ds.GetClientsByNameResponse {
	resp, err := s.clientStorage.GetClientsByName(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetClientsByName", "message", err.Error())
		return nil
	}

	return resp
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
