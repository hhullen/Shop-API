package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"shopapi/internal/clients/postgres/sqlc"
	ds "shopapi/internal/datastruct"
	"shopapi/internal/supports"
)

func (c *Client) AddClient(req *ds.AddClientRequest) (*ds.AddClientResponse, error) {
	uid := supports.GetUUIDIfEmpty(req.Uid)

	err := c.db.ExecTx(defaultTxOpt, func(ctx context.Context, q IQuerier) error {

		addresId, err := q.InsertAddress(ctx, sqlc.InsertAddressParams{
			Country: req.Address.Country,
			City:    req.Address.City,
			Street:  req.Address.Street,
		})
		if err != nil {
			return err
		}

		uid, err = q.InsertClient(ctx, sqlc.InsertClientParams{
			Uid:              uid,
			ClientName:       req.Name,
			ClientSurname:    req.Surname,
			Birthday:         time.Time(req.Birthday),
			Gender:           string(req.Gender),
			RegistrationDate: time.Time(req.RegistrationDate),
			AddressID:        addresId,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &ds.AddClientResponse{Uid: &uid}, nil
}

func (c *Client) DeleteClient(req *ds.DeleteClientRequest) (resp *ds.DeleteClientResponse, err error) {

	err = c.db.ExecTx(defaultTxOpt, func(ctx context.Context, q IQuerier) error {

		addressId, err := q.DeleteClient(ctx, req.Uid)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			resp = &ds.DeleteClientResponse{
				Status: ds.Status{Message: ds.StatusNotFound},
			}
			return nil
		}

		clients, err := q.CalculateClientsWithAddress(ctx, addressId)
		if err != nil {
			return err
		}

		suppliers, err := q.CalculateSuppliersWithAddress(ctx, addressId)
		if err != nil {
			return err
		}

		if clients == 0 && suppliers == 0 {
			err := q.DeleteAddress(ctx, addressId)
			if err != nil {
				return err
			}
		}

		resp = &ds.DeleteClientResponse{
			Status: ds.Status{Message: ds.StatusOK},
		}

		return nil
	})

	return
}

func (c *Client) GetClients(req *ds.GetClientsRequest) (*ds.GetClientsResponse, error) {
	var clients []sqlc.ClientDetail
	var err error

	q := c.db.Querier()

	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	if req.Limit == 0 && req.Offset == 0 {
		clients, err = q.GetAllClients(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		clients, err = q.GetClientsPage(ctx, sqlc.GetClientsPageParams{
			Offset: int32(req.Offset),
			Limit:  int32(req.Limit),
		})
		if err != nil {
			return nil, err
		}
	}

	resp := &ds.GetClientsResponse{}
	resp.Clients = make([]ds.Client, 0, len(clients))
	for _, c := range clients {
		resp.Clients = append(resp.Clients, ds.Client{
			Birthday:         ds.DateOnly(c.Birthday),
			RegistrationDate: ds.DateOnly(c.RegistrationDate),
			Name:             c.ClientName,
			Surname:          c.ClientSurname,
			Gender:           ds.Gender(c.Gender),
			Uid:              c.Uid,
			Address: &ds.Address{
				Country: c.Country,
				City:    c.City,
				Street:  c.Street,
			},
		})
	}

	return resp, nil
}

func (c *Client) GetClientsByName(req *ds.GetClientsByNameRequest) (*ds.GetClientsByNameResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	clients, err := c.db.Querier().GetClientsWithName(ctx, sqlc.GetClientsWithNameParams{
		ClientName:    req.Name,
		ClientSurname: req.Surname,
	})
	if err != nil {
		return nil, err
	}

	resp := &ds.GetClientsByNameResponse{}
	resp.Clients = make([]ds.Client, 0, len(clients))
	for _, c := range clients {
		resp.Clients = append(resp.Clients, ds.Client{
			Birthday:         ds.DateOnly(c.Birthday),
			RegistrationDate: ds.DateOnly(c.RegistrationDate),
			Name:             c.ClientName,
			Surname:          c.ClientSurname,
			Gender:           ds.Gender(c.Gender),
			Uid:              c.Uid,
			Address: &ds.Address{
				Country: c.Country,
				City:    c.City,
				Street:  c.Street,
			},
		})
	}

	return resp, nil
}

func (c *Client) PatchClientAddress(req *ds.PatchClientAddressRequest) (*ds.PatchClientAddressResponse, error) {
	ctx, cancel := c.db.CtxWithCancel()
	defer cancel()

	_, err := c.db.Querier().UpdateClientAddress(ctx, sqlc.UpdateClientAddressParams{
		Country: req.Address.Country,
		City:    req.Address.City,
		Street:  req.Address.Street,
		Uid:     req.Uid,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return &ds.PatchClientAddressResponse{
			Status: ds.Status{Message: ds.StatusNotFound},
		}, nil
	}

	return &ds.PatchClientAddressResponse{
		Status: ds.Status{Message: ds.StatusOK},
	}, nil
}
