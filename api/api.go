// Example - demonstrates REST API server implementation tests.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DATA-DOG/godog"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fail(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	data := struct {
		Version string `json:"version"`
	}{Version: godog.Version}

	ok(w, data)
}

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(resp))
}

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if s, ok := data.(string); ok {
		fmt.Fprintf(w, s)
		return
	}

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fail(w, "oops something evil has happened", 500)
		return
	}

	fmt.Fprintf(w, string(resp))
}