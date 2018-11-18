package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BNKServiceTestSuite struct {
	suite.Suite
}

func NewBNKServiceTestSuite() *BNKServiceTestSuite {
	return &BNKServiceTestSuite{}
}

func TestBNKServiceTestSute(t *testing.T) {
	suite.Run(t, NewBNKServiceTestSuite())
}

func (ts *BNKServiceTestSuite) SetupSuite() {
	fmt.Println("Setup")
}

func (ts *BNKServiceTestSuite) TearDownSuite() {
	fmt.Println("Teardown")
}

func (ts *BNKServiceTestSuite) TestGetBNKMembers_GivenJSON_ExpectMembers() {
	requester := NewMockRequester()
	requester.On("Get",
		"https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json").
		Return(`
		[
			{
				"birthday":"1996-05-02",
				"blood_type":"B",
				"english_first_name":"CHERPRANG",
				"english_last_name":"AREEKUL",
				"facebook":"bnk48official.cherprang",
				"height":160,
				"hobby":"Cosplay",
				"instagram":"cherprang.bnk48official",
				"like":[
					"White color",
					"Tofu"
				],
				"nickname":"Cherprang",
				"province":"Bangkok",
				"thai_first_name":"\u0e40\u0e0c\u0e2d\u0e1b\u0e23\u0e32\u0e07",
				"thai_last_name":"\u0e2d\u0e32\u0e23\u0e35\u0e22\u0e4c\u0e01\u0e38\u0e25"
			},
			{
				"birthday":"2002-06-17",
				"blood_type":"AB",
				"english_first_name":"CHRISTIN",
				"english_last_name":"LARSEN",
				"facebook":"bnk48official.namhom",
				"height":165,
				"hobby":"Playing Guitar",
				"instagram":"namhom.bnk48official",
				"like":[
					"White color",
					"Chocolate"
				],
				"nickname":"Namhom",
				"province":"Sisaket",
				"thai_first_name":"\u0e04\u0e23\u0e34\u0e2a\u0e15\u0e34\u0e19",
				"thai_last_name":"\u0e25\u0e32\u0e23\u0e4c\u0e40\u0e0b\u0e48\u0e19"
			}]`, nil)

	bnkSvc := NewBNKService()
	mems, err := bnkSvc.GetBNKMembers(requester)

	is := assert.New(ts.T())
	if is.NoError(err) {
		is.Equal(2, len(mems))
		is.Equal("CHERPRANG", mems[0].EnglishFirstname)
		is.Equal("AREEKUL", mems[0].EnglishLastname)

		is.Equal("CHRISTIN", mems[1].EnglishFirstname)
		is.Equal("LARSEN", mems[1].EnglishLastname)
	}

	requester.AssertExpectations(ts.T())
}

func (ts *BNKServiceTestSuite) TestGetBNKMembers_GivenRequestError_ExpectError() {
	requester := NewMockRequester()
	requester.On("Get",
		"https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json").
		Return("", fmt.Errorf("Mock Error"))

	bnkSvc := NewBNKService()
	mems, err := bnkSvc.GetBNKMembers(requester)

	is := assert.New(ts.T())
	if is.Nil(mems) {
		is.Equal("Error Request:Mock Error", err.Error())
	}

	requester.AssertExpectations(ts.T())
}

func (ts *BNKServiceTestSuite) TestGetBNKMembers_GivenInvalidJSON_ExpectError() {
	requester := NewMockRequester()
	requester.On("Get",
		"https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json").
		Return("<Invalid JSON>", nil)

	bnkSvc := NewBNKService()
	mems, err := bnkSvc.GetBNKMembers(requester)

	is := assert.New(ts.T())
	if is.Nil(mems) {
		is.Equal("Error JSON:invalid character '<' looking for beginning of value", err.Error())
	}

	requester.AssertExpectations(ts.T())
}
