package handlers

import (
	"net/http"
)

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "applicaiton/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	////w.Header().Add("Access-Control-Allow-Methods", "PUT")
	////w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	http.ServeFile(w, r, "swagger.json")
}
