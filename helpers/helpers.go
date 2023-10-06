// Package helpers provides utility functions and structures for a secure web API.
package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/davidandw190/coffeeshop-api-go/services"
)

// Envelope is a generic map for JSON responses.
type Envelope map[string]interface{}

// Message represents the logger configuration for the application.
type Message struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// MessageLogs contains the default logger instances.
var MessageLogs = &Message{
	InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
}

// ReadJSON reads and decodes JSON data from the request body.
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	const maxBytes = 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	// Ensure that the request body contains only one JSON object.
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have a single JSON object only")
	}

	return nil
}

// WriteJSON encodes and writes JSON response data with optional headers.
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			w.Header()[key] = val
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err = w.Write(out); err != nil {
		return err
	}

	return nil
}

// ErrorJSON responds with a JSON error message and optional status code.
func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload = services.JsonResponse{
		Error:   true,
		Message: err.Error(),
	}
	WriteJSON(w, statusCode, payload)
}
