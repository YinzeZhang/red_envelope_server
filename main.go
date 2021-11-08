package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"red_envelop_server/routers"
	"red_envelop_server/sql"
)

func main() {

	db, err := gorm.Open("mysql", "root:3306@tcp(headless-service-port-0:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	//defer db.Close()
	db.AutoMigrate(&sql.User{})
	db.AutoMigrate(&sql.Envelope{})

	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
