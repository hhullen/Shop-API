package postgres

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

var errTest = errors.New("error")

type TestClient struct {
	ctx         context.Context
	clientMock  *MockIDB
	querierMock *MockIQuerier
	client      *Client
}

func NewTestClient(t *testing.T) *TestClient {
	mc := gomock.NewController(t)

	tc := &TestClient{
		clientMock:  NewMockIDB(mc),
		querierMock: NewMockIQuerier(mc),
		ctx:         context.Background(),
	}

	tc.client = buildClient(tc.clientMock)

	return tc
}
