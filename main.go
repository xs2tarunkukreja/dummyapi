package main

import (
	"log"
	"net/http"

	"github.com/xs2tarunkukreja/dummyapi/helper"
	_ "github.com/xs2tarunkukreja/dummyapi/handlers"
)

func main() {
	helper.Routes()
	log.Println("Going to start server")
	log.Fatal(http.ListenAndServe(":3010", nil))
}
