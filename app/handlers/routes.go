package handlers

import (
	"fizzbuzz-leboncoin/app/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HandleRequest gere toute les routes du serveur HTTP
func HandleRequest(router *mux.Router, c *config.Client) {

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/fizzbuzz", fizzbuzz(c)).Methods("GET")
	router.HandleFunc("/statistics", requestNameForMaxHitNumber(c)).Methods("GET")

	// Documentation
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

}

// @Summary Index of the api
// @Success 200
// @Router / [get]
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	fmt.Fprintf(w, "<h1>Hi there, welcome to the fizzbuzz-leboncoin api !</h1>")
}
