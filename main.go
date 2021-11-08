package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"red_envelop_server/routers"
	"red_envelop_server/sql"
)

func main() {

	db, _ := gorm.Open("mysql", "root:3306@tcp(172.27.11.82:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.AutoMigrate(&sql.User{})
	db.AutoMigrate(&sql.Envelope{})

	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
