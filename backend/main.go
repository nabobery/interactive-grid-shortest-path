package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
	"grid-dfs/controllers"
)



func main() {
	fmt.Println("Grid BFS solver")
	r := mux.NewRouter()

	// The /find-path endpoint now uses the BFS handler.
	r.HandleFunc("/find-path", controllers.BFSHandler).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}