package postgres

import (
	"database/sql"
	"errors"
	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
)

func (c *Client) AddImage(req *ds.AddImageRequest) (*ds.AddImageResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	uid := supports.GetUUIDIfEmpty(req.Uid)

	uid, err := c.db.Querier().AddImage(ctx, sqlc.AddImageParams{
		Uid:   uid,
		Image: req.Image,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.AddImageResponse{
			Status: ds.Status{Message: ds.StatusAlreadyExists},
		}, nil
	}

	return &ds.AddImageResponse{
		Uid: &uid,
	}, nil
}

func (c *Client) UpdateImage(req *ds.UpdateImageRequest) (*ds.UpdateImageResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	_, err := c.db.Querier().UpdateImage(ctx, sqlc.UpdateImageParams{
		Uid:   req.Uid,
		Image: req.Image,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.UpdateImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}, nil
	}

	return &ds.UpdateImageResponse{
		Status: ds.Status{Message: ds.StatusOK},
	}, nil
}

func (c *Client) DeleteImage(req *ds.DeleteImageRequest) (*ds.DeleteImageResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	_, err := c.db.Querier().DeleteImage(ctx, req.Uid)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.DeleteImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}, nil
	}

	return &ds.DeleteImageResponse{
		Status: ds.Status{Message: ds.StatusOK},
	}, nil
}

func (c *Client) GetProductImage(req *ds.GetProductImageRequest) (*ds.GetProductImageResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	img, err := c.db.Querier().GetProductImage(ctx, req.ProductUid)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.GetProductImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}, nil
	}

	return &ds.GetProductImageResponse{
		Image: img.Image,
		Uid:   &img.Uid,
	}, nil
}

func (c *Client) GetImage(req *ds.GetImageRequest) (*ds.GetImageResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	img, err := c.db.Querier().GetImage(ctx, req.Uid)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.GetImageResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}, nil
	}

	return &ds.GetImageResponse{
		Uid:   &req.Uid,
		Image: img,
	}, nil
}
