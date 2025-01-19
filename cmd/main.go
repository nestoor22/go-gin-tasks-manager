package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-tasks-manager/routers"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	routers.AddUserRoutes(router)
	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
