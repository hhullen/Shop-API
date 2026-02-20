package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

var errTest = errors.New("error")

type TestService struct {
	loggerMock          *MockILogger
	cacheMock           *MockICache
	clientStorageMock   *MockIClientStorage
	productStorageMock  *MockIProductStorage
	supplierStorageMock *MockISupplierStorage
	imageStorageMock    *MockIImageStorage
	srv                 *Service
}

func NewTestService(t *testing.T) *TestService {
	mc := gomock.NewController(t)
	s := &TestService{
		loggerMock:          NewMockILogger(mc),
		cacheMock:           NewMockICache(mc),
		clientStorageMock:   NewMockIClientStorage(mc),
		imageStorageMock:    NewMockIImageStorage(mc),
		productStorageMock:  NewMockIProductStorage(mc),
		supplierStorageMock: NewMockISupplierStorage(mc),
	}

	s.srv = NewService(context.Background(), s.loggerMock, s.cacheMock, s.clientStorageMock,
		s.productStorageMock, s.supplierStorageMock, s.imageStorageMock)

	return s
}
