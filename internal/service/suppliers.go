package service

import (
	ds "shopapi/internal/datastruct"
)

func (s *Service) AddSupplier(req *ds.AddSupplierRequest) *ds.AddSupplierResponse {
	resp, err := s.supplierStorage.AddSupplier(req)
	if err != nil {
		s.logger.ErrorKV("failed on AddSupplier", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) UpdateSupplierAddress(req *ds.UpdateSupplierAddressRequest) *ds.UpdateSupplierAddressResponse {
	resp, err := s.supplierStorage.UpdateSupplierAddress(req)
	if err != nil {
		s.logger.ErrorKV("failed on UpdateSupplierAddress", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("UpdateSupplierAddress", resp.GetStatus())

	return resp
}

func (s *Service) DeleteSupplier(req *ds.DeleteSupplierRequest) *ds.DeleteSupplierResponse {
	resp, err := s.supplierStorage.DeleteSupplier(req)
	if err != nil {
		s.logger.ErrorKV("failed on DeleteSupplier", "message", err.Error())
		return nil
	}

	s.logHandlerStatus("DeleteSupplier", resp.GetStatus())

	return resp
}

func (s *Service) GetSuppliers(req *ds.GetSuppliersRequest) *ds.GetSuppliersResponse {
	resp, err := s.supplierStorage.GetSuppliers(req)
	if err != nil {
		s.logger.ErrorKV("failed on GetSuppliers", "message", err.Error())
		return nil
	}

	return resp
}

func (s *Service) GetSupplier(req *ds.GetSupplierRequest) (resp *ds.GetSupplierResponse) {
	var cached bool
	var err error

	key := makeCacheKey(imagesCacheKey, req.Uid.String())
	if !req.AvoidCache {
		resp = &ds.GetSupplierResponse{}
		cached, err = s.cache.Read(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed reading cache on GetSupplier", "message", err.Error())
		}
	}

	if !cached {
		resp, err = s.supplierStorage.GetSupplier(req)
		if err != nil {
			s.logger.ErrorKV("failed on GetSupplier", "message", err.Error())
			return nil
		}

		err = s.cache.Write(key, resp)
		if err != nil {
			s.logger.ErrorKV("failed writing cache on GetSupplier", "message", err.Error())
		}
		resp.Cached = false
	} else {
		resp.Cached = true
	}

	s.logHandlerStatus("GetSupplier", resp.GetStatus())

	return
}
