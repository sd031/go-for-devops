package mocks

import (
	"github.com/stretchr/testify/mock"
)

// ExternalAPIMock is a mock for the ExternalAPI
type ExternalAPIMock struct {
	mock.Mock
}

// FetchData is the mock function for ExternalAPI's FetchData
func (m *ExternalAPIMock) FetchData(param string) (string, error) {
	args := m.Called(param)
	return args.String(0), args.Error(1)
}
