package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ContentTypeJSON = "application/json"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, data interface{}, httpCode int) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		panic(err)
	}

	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(httpCode)
	w.Write(append(jsonBytes, []byte("\n")...))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	ErrorJSONResponse(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	ErrorJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
}

func ErrorJSONResponse(w http.ResponseWriter, message string, httpCode int) {
	JSONResponse(w, ErrorResponse{Message: message}, httpCode)
}
