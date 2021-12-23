package server

func (s *Server) routes() {
	s.router.HandleFunc("/add", s.HandleNewMember)
	//s.router.HandleFunc("/",s.HandleIndex)
}
