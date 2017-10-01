package handlers

import (
	"encoding/json"
	"net/http"
)

type VersionResponse struct {
	Version    string `json:"version"`
	BuildTime  string `json:"build_time"`
	CommitHash string `json:"commit_hash"`
}

// Variables to identiy the build
var (
	CommitHash string
	VersionTag string
	BuildTime  string
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {

	response := VersionResponse{
		Version:    VersionTag,
		BuildTime:  BuildTime,
		CommitHash: CommitHash,
	}

	json.NewEncoder(w).Encode(response)
	return
}
