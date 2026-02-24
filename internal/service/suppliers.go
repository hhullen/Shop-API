package service

import (
	"encoding/json"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"strconv"
)

func (s *Service) AddSupplier(req *ds.AddSupplierRequest) *ds.AddSupplierResponse {
	reqKey, err := json.Marshal(req.Supplier)
	if err != nil {
		s.logger.ErrorKV("failed making cache key", "message", err.Error(), "data", *req)
		req.AvoidCacheFlag.Flag = true
	}
	key := makeCacheKey("AddSupplier", req.Uid.String(), supports.GetHash(reqKey))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.AddSupplierResponse, error) {
		return s.supplierStorage.AddSupplier(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on AddSupplier", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) UpdateSupplierAddress(req *ds.UpdateSupplierAddressRequest) *ds.UpdateSupplierAddressResponse {
	reqKey, err := json.Marshal(req.Address)
	if err != nil {
		s.logger.ErrorKV("failed making cache key", "message", err.Error(), "data", *req)
		req.AvoidCacheFlag.Flag = true
	}
	key := makeCacheKey("UpdateSupplierAddress", req.Uid.String(), supports.GetHash(reqKey))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.UpdateSupplierAddressResponse, error) {
		return s.supplierStorage.UpdateSupplierAddress(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on UpdateSupplierAddress", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("UpdateSupplierAddress", resp.GetStatus())

	return resp
}

func (s *Service) DeleteSupplier(req *ds.DeleteSupplierRequest) *ds.DeleteSupplierResponse {
	key := makeCacheKey("DeleteSupplier", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.DeleteSupplierResponse, error) {
		return s.supplierStorage.DeleteSupplier(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on DeleteSupplier", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteSupplier", resp.GetStatus())

	return resp
}

func (s *Service) GetSuppliers(req *ds.GetSuppliersRequest) *ds.GetSuppliersResponse {
	key := makeCacheKey("GetSuppliers", strconv.FormatInt(req.Limit, 10), strconv.FormatInt(req.Offset, 10))

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetSuppliersResponse, error) {
		return s.supplierStorage.GetSuppliers(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetSuppliers", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) GetSupplier(req *ds.GetSupplierRequest) (resp *ds.GetSupplierResponse) {
	key := makeCacheKey("GetSupplier", req.Uid.String())

	resp, err := execWithCache(s, key, req.AvoidCache(), func() (*ds.GetSupplierResponse, error) {
		return s.supplierStorage.GetSupplier(req)
	})

	if err != nil {
		s.logger.ErrorKV("failed on GetSupplier", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("GetSupplier", resp.GetStatus())

	return
}
