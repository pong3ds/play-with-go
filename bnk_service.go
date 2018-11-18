package main

import (
	"encoding/json"
	"errors"

	"github.com/pong3ds/play-with-go/members"
)

// BNKService is the service for BNK
type BNKService struct{}

// NewBNKService return new BNKService
func NewBNKService() *BNKService {
	return &BNKService{}
}

// GetBNKMembers return all bnk members receive from BNK Api
func (svc *BNKService) GetBNKMembers(request IRequester) ([]*members.Member, error) {
	body, err := request.Get("https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json")
	if err != nil {
		return nil, errors.New("Error Request:" + err.Error())
	}

	mems := make([]*members.Member, 0)
	err = json.Unmarshal([]byte(body), &mems)
	if err != nil {
		return nil, errors.New("Error JSON:" + err.Error())
	}

	return mems, nil
}
