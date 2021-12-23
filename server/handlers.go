package server

import (
	"MemberClub/member"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) HandleNewMember(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	data := map[string]string{}
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}
	for credential, value := range data {
		valid := Validation(ValidationPair{
			CredToBeValidated: credential,
			Value:             value,
		})
		if !valid {
			err := s.respond(w, r, map[string]string{
				"message": "invalid input",
			}, http.StatusBadRequest)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	memberToCome := member.NewMember(data["name"], data["email"])
	err = s.userRepo.InsertNewMember(memberToCome)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Println(s.userRepo)
}
