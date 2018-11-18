package main

import "github.com/stretchr/testify/mock"

// MockRequester is mock
type MockRequester struct {
	mock.Mock
}

// NewMockRequester return new mock
func NewMockRequester() *MockRequester {
	return &MockRequester{}
}

// Get return string
func (r *MockRequester) Get(url string) (string, error) {
	args := r.Called(url)
	return args.String(0), args.Error(1)
}
