package main

import (
	"MemberClub/repo"
	"MemberClub/server"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	ctx := context.Background()
	userRepository := repo.NewRepo()
	serv := server.Init(ctx, router, userRepository)

	var addr = ":80"

	fmt.Println("Listening and serving on address", addr)
	err := http.ListenAndServe(addr, serv)
	if err != nil {
		log.Fatalln(err)
	}
}
