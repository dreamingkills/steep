package models

var Models = []interface{}{Merchant{}}

type Merchant struct {
	ID   uint32
	Name string
	URL  string
}