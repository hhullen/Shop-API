package service

import (
	ds "shopapi/internal/datastruct"
)

func (s *Service) AddProduct(req *ds.AddProductRequest) *ds.AddProductResponse {
	resp, err := s.productStorage.AddProduct(req)
	if err != nil {
		s.logger.ErrorKV("failed on AddProduct", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("AddProduct", resp.GetStatus())

	return resp
}

func (s *Service) DecreaseProducts(req *ds.DecreaseProductsRequest) *ds.DecreaseProductsResponse {
	resp, err := s.productStorage.DecreaseProducts(req)
	if err != nil {
		s.logger.ErrorKV("failed on DecreaseProducts", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DecreaseProducts", resp.GetStatus())

	return resp
}

func (s *Service) GetProduct(req *ds.GetProductRequest) (resp *ds.GetProductResponse) {
	var cached bool
	var err error

	key := makeCacheKey(imagesCacheKey, req.Uid.String())
	if !req.AvoidCache {
		resp = &ds.GetProductResponse{}
		cached, err = s.cache.Read(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed reading cache on GetProduct", "message", err.Error())
		}
	}

	if !cached {
		resp, err = s.productStorage.GetProduct(req)
		if err != nil {
			s.logger.ErrorKV("failed on GetProduct", "message", err.Error())
			return nil
		}

		err = s.cache.Write(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed writing cache on GetProduct", "message", err.Error())
		}
		resp.Cached = false
	} else {
		resp.Cached = true
	}

	s.logHandlerStatus("GetProduct", resp.GetStatus())

	return
}

func (s *Service) GetProducts(req *ds.GetProductsRequest) *ds.GetProductsResponse {
	resp, err := s.productStorage.GetProducts(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetProducts", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) DeleteProduct(req *ds.DeleteProductRequest) *ds.DeleteProductResponse {
	resp, err := s.productStorage.DeleteProduct(req)
	if err != nil {
		s.logger.ErrorKV("failed on DeleteProduct", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteProduct", resp.GetStatus())

	return resp
}
