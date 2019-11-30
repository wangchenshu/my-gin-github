package main

import (
	"log"
	"mygin/httpd/routes"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8088", routes.Engine()))
}
