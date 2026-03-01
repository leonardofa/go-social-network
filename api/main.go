package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Init()
	routers := router.Init()

	fmt.Printf("Starting GO - Social Network - API. 🚀🚀🚀: http://localhost:%d", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), routers))
}
