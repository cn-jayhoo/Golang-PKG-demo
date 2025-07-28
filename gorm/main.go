package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return
	}
	_ = db // #Gorm.DB
}

func DSN() string {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"demo",
		"demo",
		fmt.Sprintf("%s:%d", "172.18.1.194", 3306),
		"demo-schema",
	)
	return dsn
}
