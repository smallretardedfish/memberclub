package server

import (
	"MemberClub/member"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
)

func (s *Server) HandleNewMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Method", "true")

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
	if err != nil {
		log.Println(err)
		return
	}
	userTable, err := json.Marshal(s.userRepo.GetAllMembers())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(userTable))
}

func (s *Server) FetchAllMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Method", "true")
	userTable := s.userRepo.GetAllMembers()
	table, err := json.Marshal(userTable)
	res := gjson.GetBytes(table, "@values").Value()
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
}
