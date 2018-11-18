package main

import (
	"encoding/json"
	"errors"

	"github.com/pong3ds/play-with-go/members"
)

// IBNKService is interface for BNK service
type IBNKService interface {
	GetBNKMembers() ([]*members.Member, error)
}

// BNKService is the service for BNK
type BNKService struct {
	ctx IContext
}

// NewBNKService return new BNKService
func NewBNKService(ctx IContext) *BNKService {
	return &BNKService{
		ctx: ctx,
	}
}

// GetBNKMembers return all bnk members receive from BNK Api
func (svc *BNKService) GetBNKMembers() ([]*members.Member, error) {
	body, err := svc.ctx.Requester().Get("https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json")
	if err != nil {
		svc.ctx.Logger().Log(err.Error())
		return nil, errors.New("Error Request:" + err.Error())
	}

	mems := make([]*members.Member, 0)
	err = json.Unmarshal([]byte(body), &mems)
	if err != nil {
		return nil, errors.New("Error JSON:" + err.Error())
	}

	return mems, nil
}
