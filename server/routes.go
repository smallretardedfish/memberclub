package server

func (s *Server) routes() {
	s.router.HandleFunc("/member", s.HandleNewMember).Methods("POST")
	//s.router.HandleFunc("/",s.HandleIndex)
}
