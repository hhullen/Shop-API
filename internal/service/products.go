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

func (s *Service) GetProduct(req *ds.GetProductRequest) *ds.GetProductResponse {
	resp, err := s.productStorage.GetProduct(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetProduct", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DecreaseProducts", resp.GetStatus())

	return resp
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
