//Provided by the professor for using json files
package initializers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DecodeJSONRequest unmarshals a JSON from src and stores it in dst (must be a pointer).
// If the unmarshalling fails, a 400 with an error message is written on w.
func DecodeJSONRequest(dst interface{}, src io.Reader, w http.ResponseWriter) (succeeded bool) {
	if err := json.NewDecoder(src).Decode(dst); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request JSON (%v).", err)
		return false
	}
	return true
}

// EncodeJSONResponse marshals an object to JSON and sends a 200 with the
// object, or sends a 500 with an error message if marshalling fails.
func EncodeJSONResponse(w http.ResponseWriter, response interface{}) {
	if raw, err := json.Marshal(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode query result to JSON (%v).", err)
	} else {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(raw)
	}
}
