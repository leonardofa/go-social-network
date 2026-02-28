package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	fmt.Println("Starting GO - Social Network - API. 🚀🚀🚀")

	r := router.Execute()
	log.Fatal(http.ListenAndServe(":5000", r))
}
