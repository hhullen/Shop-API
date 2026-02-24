package postgres

import (
	"context"
	"database/sql"
	"errors"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
)

func (c *Client) AddSupplier(req *ds.AddSupplierRequest) (resp *ds.AddSupplierResponse, err error) {
	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, qtx IQuerier) error {
		addresId, err := qtx.InsertAddress(ctx, sqlc.InsertAddressParams{
			Country: req.Address.Country,
			City:    req.Address.City,
			Street:  req.Address.Street,
		})
		if err != nil {
			return err
		}

		uid, err := qtx.InsertSupplier(ctx, sqlc.InsertSupplierParams{
			Uid:         supports.GetUUIDIfEmpty(req.Uid),
			Name:        req.Name,
			PhoneNumber: string(req.PhoneNumber),
			AddressID:   addresId,
		})
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			resp = &ds.AddSupplierResponse{
				Status: ds.Status{Message: ds.StatusAlreadyExists},
			}
			return nil
		}
		resp = &ds.AddSupplierResponse{Uid: &uid}

		return nil
	})

	return
}

func (c *Client) UpdateSupplierAddress(req *ds.UpdateSupplierAddressRequest) (resp *ds.UpdateSupplierAddressResponse, err error) {

	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, qtx IQuerier) error {
		addressId, err := qtx.InsertAddress(ctx, sqlc.InsertAddressParams{
			Country: req.Address.Country,
			City:    req.Address.City,
			Street:  req.Address.Street,
		})
		if err != nil {
			return err
		}

		_, err = qtx.UpdateSupplierAddress(ctx, sqlc.UpdateSupplierAddressParams{
			Uid:       req.Uid,
			AddressID: addressId,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				var suppliers int64
				suppliers, err = qtx.CalculateSuppliersWithAddress(ctx, addressId)
				if err != nil {
					return err
				}

				var clients int64
				clients, err = qtx.CalculateClientsWithAddress(ctx, addressId)
				if err != nil {
					return err
				}

				if suppliers == 0 && clients == 0 {
					err = qtx.DeleteAddress(ctx, addressId)
					if err != nil {
						return err
					}
				}

				resp = &ds.UpdateSupplierAddressResponse{
					Status: ds.Status{Message: ds.StatusNotFound},
				}
				return nil
			}
			return err
		}

		resp = &ds.UpdateSupplierAddressResponse{
			Status: ds.Status{Message: ds.StatusOK},
		}

		return nil
	})

	return
}

func (c *Client) DeleteSupplier(req *ds.DeleteSupplierRequest) (resp *ds.DeleteSupplierResponse, err error) {

	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, qtx IQuerier) error {
		addressId, err := qtx.DeleteSupplier(ctx, req.Uid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				resp = &ds.DeleteSupplierResponse{
					Status: ds.Status{Message: ds.StatusNotFound},
				}
				return nil
			}
			return err
		}

		suppliers, err := qtx.CalculateSuppliersWithAddress(ctx, addressId)
		if err != nil {
			return err
		}

		clients, err := qtx.CalculateClientsWithAddress(ctx, addressId)
		if err != nil {
			return err
		}

		if suppliers == 0 && clients == 0 {
			err := qtx.DeleteAddress(ctx, addressId)
			if err != nil {
				return err
			}
		}

		resp = &ds.DeleteSupplierResponse{Status: ds.Status{Message: ds.StatusOK}}
		return nil
	})

	return
}

func (c *Client) GetSuppliers(req *ds.GetSuppliersRequest) (*ds.GetSuppliersResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	var err error
	var suppliers []sqlc.SupplierDetail
	if req.Limit == 0 && req.Offset == 0 {
		suppliers, err = c.db.Querier().GetAllSuppliers(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		suppliers, err = c.db.Querier().GetSuppliersPage(ctx, sqlc.GetSuppliersPageParams{
			Offset: int32(req.Offset),
			Limit:  int32(req.Limit),
		})
		if err != nil {
			return nil, err
		}
	}

	resp := &ds.GetSuppliersResponse{
		Suppliers: make([]ds.Supplier, len(suppliers)),
	}
	for i := range suppliers {
		s := &suppliers[i]
		resp.Suppliers[i] = ds.Supplier{
			Uid:         s.Uid,
			Name:        s.Name,
			PhoneNumber: ds.PhoneNumber(s.PhoneNumber),
			Address: &ds.Address{
				Country: s.Country,
				City:    s.City,
				Street:  s.Street,
			},
		}
	}

	return resp, nil
}

func (c *Client) GetSupplier(req *ds.GetSupplierRequest) (*ds.GetSupplierResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	s, err := c.db.Querier().GetSupplier(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ds.GetSupplierResponse{
				Status: ds.Status{Message: ds.StatusNotFound},
			}, nil
		} else {
			return nil, err
		}
	}

	return &ds.GetSupplierResponse{
		Supplier: &ds.Supplier{
			Uid:         s.Uid,
			Name:        s.Name,
			PhoneNumber: ds.PhoneNumber(s.PhoneNumber),
			Address: &ds.Address{
				Country: s.Country,
				City:    s.City,
				Street:  s.Street,
			},
		},
	}, nil
}
