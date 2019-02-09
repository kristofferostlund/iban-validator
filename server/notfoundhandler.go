package server

import (
	"net/http"

	"github.com/kristofferostlund/pfc-iban-validator/server/responses"
)

func (s *Server) notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.ErrorJSONResponse(w, "Not found", http.StatusNotFound)
	}
}
