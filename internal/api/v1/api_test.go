package api

import (
	"context"
	"shopapi/internal/service"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

type TestAPI struct {
	clientMock     *MockIClientService
	imageMock      *MockIImageService
	productMock    *MockIProductService
	supplierMock   *MockISupplierService
	serverMock     *MockIServer
	routerMock     *MockIRouter
	loggerMock     *service.MockILogger
	responseWriter *MockResponseWriter
	api            *API
}

func NewTestApi(ctx context.Context, t *testing.T) *TestAPI {
	mc := gomock.NewController(t)
	ta := &TestAPI{
		clientMock:     NewMockIClientService(mc),
		imageMock:      NewMockIImageService(mc),
		productMock:    NewMockIProductService(mc),
		supplierMock:   NewMockISupplierService(mc),
		serverMock:     NewMockIServer(mc),
		routerMock:     NewMockIRouter(mc),
		loggerMock:     service.NewMockILogger(mc),
		responseWriter: NewMockResponseWriter(mc),
	}

	ta.routerMock.EXPECT().HandleFunc(gomock.Any(), gomock.Any()).MinTimes(1)

	ta.api = buildAPI(ctx, ta.loggerMock, ta.serverMock, ta.routerMock,
		ta.clientMock, ta.productMock, ta.supplierMock, ta.imageMock)

	return ta
}

func TestBuildApi(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	NewTestApi(ctx, t)
}
