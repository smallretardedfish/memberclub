package main

import (
	"MemberClub/repo"
	"MemberClub/server"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	ctx := context.Background()
	userRepository := repo.NewRepo()
	serv := server.Init(ctx, router, userRepository)

	var addr = ":8080"
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET"},
		AllowCredentials: true,
	})
	handler := c.Handler(serv)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	fmt.Printf("staring web server on %s\n", addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
