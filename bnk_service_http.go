package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// BNKServiceHTTP is http handler for BNK service
type BNKServiceHTTP struct{}

// NewBNKServiceHTTP return new BNKServiceHTTP
func NewBNKServiceHTTP() *BNKServiceHTTP {
	return &BNKServiceHTTP{}
}

// RegisterServices register all services for bnk http service
func (svc *BNKServiceHTTP) RegisterServices(e *echo.Echo) error {
	e.GET("/bnk48/members", svc.handleBNK48Members)
	e.GET("/bnk48/senbetsu", svc.handleBNK48Senbetsu)
	return nil
}

func (svc *BNKServiceHTTP) handleBNK48Senbetsu(c echo.Context) error {
	return nil
}

func (svc *BNKServiceHTTP) handleBNK48Members(c echo.Context) error {
	bnkSvc := NewBNKService()
	members, err := bnkSvc.GetBNKMembers(NewRequester())
	if err != nil {
		c.Error(err)
	}
	return c.JSONPretty(http.StatusOK, members, "  ")
}
