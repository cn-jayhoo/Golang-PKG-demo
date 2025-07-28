package model

import "github.com/shopspring/decimal"

type TPrice struct {
	UID   int             `gorm:"primary_key;type:int(11) unsigned auto_increment;column:uid"`
	Price decimal.Decimal `gorm:"type:decimal(10,2) unsigned;"`
}

func (t *TPrice) TableName() string {
	return "t_decimal"
}
