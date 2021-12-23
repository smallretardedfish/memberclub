package server

import (
	"MemberClub/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {
	router   *mux.Router
	userRepo *repo.UserRepo
}

func Init(router *mux.Router, userRepo *repo.UserRepo) *Server {
	s := &Server{
		router:   router,
		userRepo: userRepo,
	}
	s.routes()
	return s
}
func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) error {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err, time.Now())
		return err
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
