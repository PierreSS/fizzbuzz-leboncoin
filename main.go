//go:generate swag init

package main

import (
	"fizzbuzz-leboncoin/app/config"
	"fizzbuzz-leboncoin/app/handlers"
	"log"
	"net/http"
	"os"

	_ "fizzbuzz-leboncoin/docs"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Set port
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port + ".")
	}
	return ":" + port
}

// @title FizzBuzz LeBonCoin API
// @version 1
// @description This is the documentation for the golang api of fizzbuzz-leboncoin

// @contact.name API Support
// @contact.email pierre.saintsorny@gmail.com

// @BasePath /
func main() {
	// Init the db
	c, err := config.InitClient()
	if err != nil {
		log.Println("Client inititation failed : " + err.Error())
		return
	}
	defer c.DB.Close()

	// Configure new router and handle request
	r := mux.NewRouter().StrictSlash(true)
	handler := cors.Default().Handler(r)

	handlers.HandleRequest(r, c)

	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(getPort(), handler))
}
