package postgres

import (
	"context"
	"database/sql"
	"errors"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
	"time"
)

func (c *Client) AddProduct(req *ds.AddProductRequest) (resp *ds.AddProductResponse, err error) {

	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, qtx IQuerier) error {
		exists, err := qtx.IsImageAndSupplierExists(ctx, sqlc.IsImageAndSupplierExistsParams{
			ImageUid:    req.ImageUid,
			SupplierUid: req.SupplierUid,
		})
		if err != nil {
			return err
		}

		if !exists {
			resp = &ds.AddProductResponse{
				Status: ds.Status{Message: ds.StatusAddProductWithNoImageOrSupplier},
			}
			return nil
		}

		lastUpdate := supports.GetNowIfZero(time.Time(req.LastUpdateDate))

		uid, err := qtx.InsertProduct(ctx, sqlc.InsertProductParams{
			Uid:            supports.GetUUIDIfEmpty(req.Uid),
			Name:           req.Name,
			Category:       req.Category,
			Price:          toDBPrice(req.Price),
			AvailableStock: req.AvaliableStocks,
			LastUpdateDate: lastUpdate,
			SupplierID:     req.SupplierUid,
			ImageID:        req.ImageUid,
		})
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			resp = &ds.AddProductResponse{
				Status: ds.Status{Message: ds.StatusAlreadyExists},
			}
			return nil
		}

		resp = &ds.AddProductResponse{Uid: &uid}

		return nil
	})

	return
}

func (c *Client) DecreaseProducts(req *ds.DecreaseProductsRequest) (resp *ds.DecreaseProductsResponse, err error) {

	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, qtx IQuerier) error {
		left, err := qtx.LockStockForUpdate(ctx, req.Uid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				resp = &ds.DecreaseProductsResponse{
					Status: ds.Status{Message: ds.StatusNotFound},
				}
				return nil
			} else {
				return err
			}
		}

		if left < req.Amount {
			resp = &ds.DecreaseProductsResponse{
				Status: ds.Status{Message: ds.StatusDecreaseProductsFailed},
				Left:   &left,
			}
			return nil
		}

		left, err = qtx.DecreaseProduct(ctx, sqlc.DecreaseProductParams{
			Amount: req.Amount,
			Uid:    req.Uid,
		})
		if err != nil {
			return err
		}

		resp = &ds.DecreaseProductsResponse{
			Left: &left,
		}

		return nil
	})

	return
}

func (c *Client) GetProduct(req *ds.GetProductRequest) (*ds.GetProductResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	res, err := c.db.Querier().GetProduct(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ds.GetProductResponse{
				Status: ds.Status{Message: ds.StatusNotFound},
			}, nil
		} else {
			return nil, err
		}
	}

	return &ds.GetProductResponse{
		Product: &ds.Product{
			Uid:             res.Uid,
			SupplierUid:     res.SupplierID,
			ImageUid:        res.ImageID,
			LastUpdateDate:  ds.DateOnly(res.LastUpdateDate),
			Name:            res.Name,
			Category:        res.Category,
			Price:           fromDBPrice(res.Price),
			AvaliableStocks: res.AvailableStock,
		},
	}, nil
}

func (c *Client) GetProducts(req *ds.GetProductsRequest) (*ds.GetProductsResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	var err error
	var products []sqlc.Product
	if req.Limit == 0 && req.Offset == 0 {
		products, err = c.db.Querier().GetAllProducts(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		products, err = c.db.Querier().GetProductsPage(ctx, sqlc.GetProductsPageParams{
			Offset: int32(req.Offset),
			Limit:  int32(req.Limit),
		})
		if err != nil {
			return nil, err
		}
	}

	resp := &ds.GetProductsResponse{
		Products: make([]ds.Product, len(products)),
	}
	for i := range products {
		p := &products[i]
		resp.Products[i] = ds.Product{
			Uid:             p.Uid,
			SupplierUid:     p.SupplierID,
			ImageUid:        p.ImageID,
			LastUpdateDate:  ds.DateOnly(p.LastUpdateDate),
			Name:            p.Name,
			Category:        p.Category,
			Price:           fromDBPrice(p.Price),
			AvaliableStocks: p.AvailableStock,
		}
	}

	return resp, nil
}

func (c *Client) DeleteProduct(req *ds.DeleteProductRequest) (*ds.DeleteProductResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	_, err := c.db.Querier().DeleteProduct(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ds.DeleteProductResponse{
				Status: ds.Status{Message: ds.StatusNotFound},
			}, nil
		} else {
			return nil, err
		}
	}

	return &ds.DeleteProductResponse{Status: ds.Status{Message: ds.StatusOK}}, nil
}
