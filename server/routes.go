package server

import (
	"log"
	"net/http"

	"github.com/kristofferostlund/pfc-iban-validator/server/requests"
)

func (s *Server) initRoutes() {
	log.Println("Initializing routes")

	s.router.HandleFunc("/health", s.handleHealth())
	s.router.HandleFunc("/iban/validate", requests.RouteHandler(requests.FailableHandlerMap{
		http.MethodPost: s.handleIbanValidatePost(),
	}))

	s.router.HandleFunc("/", s.notFoundHandler())
}
