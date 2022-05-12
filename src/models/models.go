package models

import "database/sql/driver"

type Merchant struct {
	ID           uint
	Name         string
	URL          string
	MerchantTeas []Tea
}

type Tea struct {
	ID         uint
	Name       string
	Type       TeaType
	MerchantID uint
	Merchant   Merchant
}

type TeaType string

const (
	black  TeaType = "black"
	green  TeaType = "green"
	oolong TeaType = "oolong"
	white  TeaType = "white"
	puerh  TeaType = "puerh"
	yellow TeaType = "yellow"
	other  TeaType = "other"
)

func (tt *TeaType) Scan(value interface{}) error {
	*tt = TeaType(value.([]byte))
	return nil
}

func (tt TeaType) Value() (driver.Value, error) {
	return string(tt), nil
}
