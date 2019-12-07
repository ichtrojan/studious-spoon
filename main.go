package main

import (
	"github.com/ichtrojan/studious-spoon/routes"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":9000", routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}
