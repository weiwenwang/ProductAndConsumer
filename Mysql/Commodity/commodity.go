package Commodity

import (
	"github.com/weiwenwang/ProductAndConsumer/Mysql"
	Commodity2 "github.com/weiwenwang/ProductAndConsumer/Commodity"
)

type Commodity struct {
	PKID    int    `gorm:"primary_key;column:PKID"`
	Name    string `gorm:"column:Name"`
	Barcode int    `gorm:"column:Barcode"`
}

var count int

const page_size = 30

func (Commodity) TableName() string {
	return "WCC_Commodity_01"
}

func Info() ([]Commodity2.Commodity, bool) {
	commodities := make([]Commodity, 0)
	Mysql.MyDb.Lock.Lock()
	Mysql.MyDb.DB.Limit(page_size).Order("PKID desc").Offset(page_size * count).Find(&commodities)
	count++
	Mysql.MyDb.Lock.Unlock()
	if (len(commodities) == 0) {
		return nil, false
	}
	node := make([]Commodity2.Commodity, 0)
	for _, v := range commodities {
		n := Commodity2.Commodity{
			v.PKID,
			v.Barcode,
		}
		node = append(node, n)
	}
	return node, true
}
