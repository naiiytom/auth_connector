package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/naiiytom/auth_connector/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting auth REST service ...")
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.HandleFunc("/auth", middleware.Authenticate).Methods("POST")

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":5000", handler))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
}
