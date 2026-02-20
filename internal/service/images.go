package service

import (
	ds "shopapi/internal/datastruct"
)

const (
	imagesCacheKey = "images"
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

func (s *Service) GetProductImage(req *ds.GetProductImageRequest) (resp *ds.GetProductImageResponse) {
	var cached bool
	var err error

	key := makeCacheKey(imagesCacheKey, req.ProductUid.String())
	if !req.AvoidCache {
		resp = &ds.GetProductImageResponse{}
		cached, err = s.cache.Read(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed reading cache on GetProductImage", "message", err.Error())
		}
	}

	if !cached {
		resp, err = s.imageStorage.GetProductImage(req)
		if err != nil {
			s.logger.ErrorKV("failed on GetProductImage", "message", err.Error())
			return nil
		}

		err = s.cache.Write(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed writing cache on GetProductImage", "message", err.Error())
		}
		resp.Cached = false
	} else {
		resp.Cached = true
	}

	s.logHandlerStatus("GetProductImage", resp.GetStatus())

	return
}

func (s *Service) GetImage(req *ds.GetImageRequest) (resp *ds.GetImageResponse) {
	var cached bool
	var err error

	key := makeCacheKey(imagesCacheKey, req.Uid.String())
	if !req.AvoidCache {
		resp = &ds.GetImageResponse{}
		cached, err = s.cache.Read(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed reading cache on GetImage", "message", err.Error())
		}
	}

	if !cached {
		resp, err = s.imageStorage.GetImage(req)
		if err != nil {
			s.logger.ErrorKV("failed on GetImage", "message", err.Error())
			return nil
		}

		err = s.cache.Write(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed writing cache on GetImage", "message", err.Error())
		}
		resp.Cached = false
	} else {
		resp.Cached = true
	}

	s.logHandlerStatus("GetImage", resp.GetStatus())

	return
}
