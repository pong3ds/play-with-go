package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/3dsinteractive/pam4-go/api/errors"
	"github.com/parnurzeal/gorequest"
	"github.com/pong3ds/play-with-go/members"
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
		return body, errors.New(fmt.Sprintf("HTTP Error code=%d", resp.StatusCode))
	}

	return body, nil
}

func getBNKMembers(request IRequester) ([]*members.Member, error) {
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

func main() {
	getBNKMembers(NewRequester())
}
