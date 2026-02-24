package service

import (
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
)

func (s *Service) AddImage(req *ds.AddImageRequest) *ds.AddImageResponse {
	key := makeCacheKey("AddImage", req.Uid.String(), supports.GetHash(req.Image))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.AddImageResponse, error) {
		return s.imageStorage.AddImage(req)
	})
	if err != nil {
		s.logger.ErrorKV("failed on AddImage", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) UpdateImage(req *ds.UpdateImageRequest) *ds.UpdateImageResponse {
	key := makeCacheKey("UpdateImage", req.Uid.String(), supports.GetHash(req.Image))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.UpdateImageResponse, error) {
		return s.imageStorage.UpdateImage(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on UpdateImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("UpdateImage", resp.GetStatus())

	return resp
}

func (s *Service) DeleteImage(req *ds.DeleteImageRequest) *ds.DeleteImageResponse {
	key := makeCacheKey("DeleteImage", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.DeleteImageResponse, error) {
		return s.imageStorage.DeleteImage(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on DeleteImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteImage", resp.GetStatus())

	return resp
}

func (s *Service) GetProductImage(req *ds.GetProductImageRequest) (resp *ds.GetProductImageResponse) {
	key := makeCacheKey("GetProductImage", req.ProductUid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetProductImageResponse, error) {
		return s.imageStorage.GetProductImage(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetProductImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetProductImage", resp.GetStatus())

	return
}

func (s *Service) GetImage(req *ds.GetImageRequest) (resp *ds.GetImageResponse) {
	key := makeCacheKey("GetImage", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetImageResponse, error) {
		return s.imageStorage.GetImage(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetImage", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetImage", resp.GetStatus())

	return
}
