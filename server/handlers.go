package server

import (
	"MemberClub/member"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}

}
func (s *Server) HandleNewMember(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

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
			return
		}

	}
	memberToCome := member.NewMember(data["name"], data["email"])
	err = s.userRepo.InsertNewMember(data["email"], memberToCome)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	err = tmpl.Execute(w, s.userRepo)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s.userRepo, s.userRepo.Size())
}
