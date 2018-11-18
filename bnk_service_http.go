package main

import (
	"net/http"
)

// BNKServiceHTTP is http handler for BNK service
type BNKServiceHTTP struct{}

// NewBNKServiceHTTP return new BNKServiceHTTP
func NewBNKServiceHTTP() *BNKServiceHTTP {
	return &BNKServiceHTTP{}
}

// RegisterServices register all services for bnk http service
func (svc *BNKServiceHTTP) RegisterServices(e *Engine) error {
	e.GET("/bnk48/members", func(ctx IContext) error {
		bnkSvc := NewBNKService(ctx)
		return svc.handleBNK48Members(ctx, bnkSvc)
	})

	e.GET("/bnk48/senbetsu", func(ctx IContext) error {
		bnkSvc := NewBNKService(ctx)
		return svc.handleBNK48Senbetsu(ctx, bnkSvc)
	})
	return nil
}

func (svc *BNKServiceHTTP) handleBNK48Senbetsu(ctx IContext, bnkSvc IBNKService) error {
	return nil
}

func (svc *BNKServiceHTTP) handleBNK48Members(ctx IContext, bnkSvc IBNKService) error {
	members, err := bnkSvc.GetBNKMembers()
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, members, "  ")
}
