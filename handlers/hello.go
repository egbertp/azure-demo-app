package handlers

import (
	"encoding/json"
	"net/http"
)

// A successResponse is a response that returns a simple JSON message
// swagger:response successResponse
type HelloResponse struct {
	// A short message
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET / hello listHello
	//
	// Returns simple JSON message
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: successResponse

	//w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers",
	//	"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//
	//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	response := HelloResponse{
		Message: "Hello world",
	}

	json.NewEncoder(w).Encode(response)
	return
}
