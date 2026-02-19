package service

import (
	ds "shopapi/internal/datastruct"
)

func (s *Service) AddImage(req *ds.AddImageRequest) *ds.AddImageResponse {
	resp, err := s.imageStorage.AddImage(req)
	if err != nil {
		s.logger.ErrorKV("failed on AddImage", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) UpdateImage(req *ds.UpdateImageRequest) *ds.UpdateImageResponse {
	resp, err := s.imageStorage.UpdateImage(req)
	if err != nil {
		s.logger.ErrorKV("failed on UpdateImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("UpdateImage", resp.GetStatus())

	return resp
}

func (s *Service) DeleteImage(req *ds.DeleteImageRequest) *ds.DeleteImageResponse {
	resp, err := s.imageStorage.DeleteImage(req)
	if err != nil {
		s.logger.ErrorKV("failed on DeleteImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteImage", resp.GetStatus())

	return resp
}

func (s *Service) GetProductImage(req *ds.GetProductImageRequest) *ds.GetProductImageResponse {
	resp, err := s.imageStorage.GetProductImage(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetProductImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetProductImage", resp.GetStatus())

	return resp
}

func (s *Service) GetImage(req *ds.GetImageRequest) *ds.GetImageResponse {
	resp, err := s.imageStorage.GetImage(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetImage", resp.GetStatus())

	return resp
}
