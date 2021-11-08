package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"red_envelop_server/routers"
	"red_envelop_server/sql"
)

func main() {

	db, _ := gorm.Open("mysql", "root:123456@tcp(headless-service-port-0:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&sql.User{})
	db.AutoMigrate(&sql.Envelope{})
	db.Close()

	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
