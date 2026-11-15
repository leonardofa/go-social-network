package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/util"
)

func main() {
	fmt.Println("Starting Social Network - WEB")
	util.LoadHtmlTemplates()
	routers := router.Init()
	fmt.Printf("Web server started. 🌎🌎🌎: http://localhost:%d", 80)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), routers))
}
