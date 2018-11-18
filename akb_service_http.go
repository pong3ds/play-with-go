package main

// AKBServiceHTTP is http service for AKB
type AKBServiceHTTP struct{}

// NewAKBServiceHTTP return new AKBServiceHTTP
func NewAKBServiceHTTP() *AKBServiceHTTP {
	return &AKBServiceHTTP{}
}

// RegisterServices register all AKB services
func (svc *AKBServiceHTTP) RegisterServices(e *Engine) error {
	return nil
}
