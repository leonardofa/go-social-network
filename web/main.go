package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	routers := router.Init()
	fmt.Printf("Starting Social Network - WEB. 🌎🌎🌎: http://localhost:%d", 80)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), routers))
}
