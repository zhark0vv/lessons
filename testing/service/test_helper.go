package service

import (
	"testing"

	"lessons/testing/service/mocks"
)

type TestHelper struct {
	DataProvider *mocks.DataProvider
	Service      *Service
}

func NewTestHelper(t *testing.T) *TestHelper {
	dp := mocks.NewDataProvider(t)

	return &TestHelper{
		DataProvider: dp,
		Service:      New(dp),
	}
}
