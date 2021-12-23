package main

import (
	"MemberClub/repo"
	"MemberClub/server"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	userRepository := repo.NewRepo()
	serv := server.Init(router, userRepository)

	var addr = ":80"

	fmt.Println("Listening and serving on address", addr)
	http.ListenAndServe(addr, serv)
}
