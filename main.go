package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"red_envelop_server/routers"
	"red_envelop_server/sql"
)

func main() {

	db, err := sql.InitDB()
	if err != nil {
		log.Println("database connection failure")
	}
	defer db.Close()

	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
