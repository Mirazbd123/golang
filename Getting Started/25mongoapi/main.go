package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mirazbd123/mongoapi/router"
)

func main() {
	fmt.Println("Mongo api !")
	r := router.Router()
	fmt.Println("Server is getting started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
