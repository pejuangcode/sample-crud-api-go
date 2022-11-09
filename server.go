package main

import (
	"log"
	"myapp/config"
	"myapp/router"
	"myapp/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	service.SetDB(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routers := gin.New()
	routers.Use(
		gin.Recovery(),
	)
	router.WebRouter(routers)

	log.Println("Listen and serve at http://localhost:" + port)
	routers.Run(":" + port)

}
