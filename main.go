// Package classification ToDo App API.
//
// this simple ToDO app REST api application serves as a learning project
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:7000
//     BasePath:
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Egbert Pot<egbert@egbertpot.nl> http://egbertpot.nl
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     - api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
//     Extensions:
//     ---
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//     ---
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"gitlab.com/egbertp/azure-demo-app/handlers"
	"gitlab.com/egbertp/azure-demo-app/helpers"
	"gitlab.com/egbertp/azure-demo-app/middleware"
)

// Variables to identiy the build
var (
	CommitHash string
	VersionTag string
	BuildTime  string
)

type Configuration struct {
	Debug   bool   `default:"false"`
	Address string `default:"0.0.0.0:7000"`
}

func main() {
	var c Configuration
	err := envconfig.Process("azure_demo_app", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// var (
	// 	httpAddr = flag.String("http", "0.0.0.0:7000", "HTTP server address")
	// )
	// flag.Parse()

	log.Println("Starting Azure Demo App")
	log.Printf("The version is: %s; the commit hash is: %s. Build time is: %s", VersionTag, CommitHash, helpers.ParseBuildTime(BuildTime).Format(time.RFC1123))
	log.Printf("listening on address %s", c.Address)

	handlers.CommitHash = CommitHash
	handlers.VersionTag = VersionTag
	handlers.BuildTime = BuildTime

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HelloHandler)
	r.HandleFunc("/version", handlers.VersionHandler)
	r.HandleFunc("/swagger", handlers.SwaggerHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(c.Address, middleware.CombinedMiddleware(r)))

}

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
