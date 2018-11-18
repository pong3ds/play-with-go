package main

import (
	"github.com/labstack/echo"
)

// AKBServiceHTTP is http service for AKB
type AKBServiceHTTP struct{}

// NewAKBServiceHTTP return new AKBServiceHTTP
func NewAKBServiceHTTP() *AKBServiceHTTP {
	return &AKBServiceHTTP{}
}

// RegisterServices register all AKB services
func (svc *AKBServiceHTTP) RegisterServices(e *echo.Echo) error {
	return nil
}
