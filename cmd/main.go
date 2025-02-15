package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-gin-tasks-manager/db"
	"go-gin-tasks-manager/routers"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	databaseConnection := db.DatabaseConnection()
	log.Println(databaseConnection)
	routers.AddUserRoutes(router)
	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
