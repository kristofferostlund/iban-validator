package server

import (
	"net/http"

	"github.com/kristofferostlund/pfc-iban-validator/iban"
	"github.com/kristofferostlund/pfc-iban-validator/server/requests"
	"github.com/kristofferostlund/pfc-iban-validator/server/responses"
)

func (s *Server) handleIbanValidatePost() requests.FailableHandler {
	type request struct {
		IBAN string `json:"iban"`
	}

	type response struct {
		IsValid bool   `json:"isValid"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		input := request{}
		if err := requests.FromJSONBody(r.Body, &input); err != nil {
			return err
		}

		if input.IBAN == "" {
			responses.ErrorJSONResponse(w, "\"iban\" is a required field", http.StatusBadRequest)
			return nil
		}

		isValid, message, err := iban.Validate(input.IBAN)
		if err != nil {
			return err
		}

		responses.JSONResponse(w, response{isValid, message}, http.StatusOK)

		return nil
	}
}
