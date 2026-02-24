package service

import (
	"encoding/json"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"strconv"
)

func (s *Service) AddProduct(req *ds.AddProductRequest) *ds.AddProductResponse {
	reqKey, err := json.Marshal(req.Product)
	if err != nil {
		s.logger.ErrorKV("failed making cache key", "message", err.Error(), "data", *req)
		req.AvoidCacheFlag.Flag = true
	}
	key := makeCacheKey("AddProduct", req.Uid.String(), supports.GetHash(reqKey))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.AddProductResponse, error) {
		return s.productStorage.AddProduct(req)
	})

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
	key := makeCacheKey("GetProduct", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetProductResponse, error) {
		return s.productStorage.GetProduct(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetProduct", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetProduct", resp.GetStatus())

	return
}

func (s *Service) GetProducts(req *ds.GetProductsRequest) *ds.GetProductsResponse {
	key := makeCacheKey("GetProducts", strconv.FormatInt(req.Limit, 10), strconv.FormatInt(req.Offset, 10))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetProductsResponse, error) {
		return s.productStorage.GetProducts(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetProducts", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) DeleteProduct(req *ds.DeleteProductRequest) *ds.DeleteProductResponse {
	key := makeCacheKey("DeleteProduct", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.DeleteProductResponse, error) {
		return s.productStorage.DeleteProduct(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on DeleteProduct", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteProduct", resp.GetStatus())

	return resp
}
