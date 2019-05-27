package Mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var MyDb struct {
	DB   *gorm.DB
	Lock sync.Mutex
}

func InitDB() () {
	//db, err := gorm.Open("mysql", "root:WOAImama188@(127.0.0.1:3306)/thomas?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "gcore:gcore@(192.168.1.6:3306)/gcore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	MyDb.DB = db
}
