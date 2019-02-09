package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kristofferostlund/iban-validator/server/responses"
)

type FailableHandler func(http.ResponseWriter, *http.Request) error
type FailableHandlerMap map[string]FailableHandler

func FromJSONBody(readable io.Reader, out interface{}) error {
	body, err := ioutil.ReadAll(readable)
	if err != nil {
		return fmt.Errorf("Failed to read request body: %v", err)
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return fmt.Errorf("Failed to unmarshal request body: %v", err)
	}

	return nil
}

func RouteHandler(handlers FailableHandlerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, exists := handlers[r.Method]
		if !exists {
			responses.MethodNotAllowed(w, r)
			return
		}

		if err := handler(w, r); err != nil {
			log.Printf("Unhandled internal server error: %s", err)
			responses.InternalServerError(w, r)
			return
		}
	}
}
