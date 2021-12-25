package server

import (
	"MemberClub/repo"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {
	ctx      context.Context
	router   *mux.Router
	userRepo repo.UserRepo
}

func Init(ctx context.Context, router *mux.Router, userRepo repo.UserRepo) *Server {
	s := &Server{
		ctx:      ctx,
		router:   router,
		userRepo: userRepo,
	}
	s.routes()
	return s
}
func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err, time.Now())
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
