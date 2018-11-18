package main

import (
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// IRequester is the interface for requester
type IRequester interface {
	Get(url string) (string, error)
}

// Requester is the object for requester
type Requester struct{}

// NewRequester return new requester
func NewRequester() *Requester {
	return &Requester{}
}

// Get the result of input url
func (r *Requester) Get(url string) (string, error) {
	resp, body, errs := gorequest.New().
		Get(url).
		End()

	if len(errs) > 0 {
		return "", errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("HTTP Error code=%d", resp.StatusCode)
	}

	return body, nil
}
