package main

import (
	"fmt"
)

func main() {
	bnkSvc := NewBNKService()
	mems, err := bnkSvc.GetBNKMembers(NewRequester())
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	for _, member := range mems {
		fmt.Println(member.Nickname)
	}
}
