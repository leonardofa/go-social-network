package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	fmt.Println("Starting Social Network - API")
	config.Init()
	routers := router.Init()

	fmt.Printf("API ready to go. 🚀🚀🚀: http://localhost:%d", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), routers))
}
