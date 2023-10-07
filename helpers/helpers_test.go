package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReadJSON(t *testing.T) {
	t.Parallel()

	t.Run("Valid JSON", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{}

		jsonData := `{"name": "John"}`

		request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		err = ReadJSON(w, request, &data)

		if err != nil {
			t.Errorf("ReadJSON() error = %v, want nil", err)
		}

		if data.Name != "John" {
			t.Errorf("ReadJSON() data.Name = %s, want John", data.Name)
		}

	})

	t.Run("Invalid JSON", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{}

		jsonData := `{"name": "John",}` // invalid JSON with extra comma

		request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		err = ReadJSON(w, request, &data)

		if err == nil {
			t.Error("ReadJSON() expected an error, got nil")
		}
	})

	t.Run("Request Body Too Large", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{}

		jsonData := `{"name": "John"}`

		request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		request.ContentLength = 1048577 // Exceeds the maxBytes limit

		w := httptest.NewRecorder()

		err = ReadJSON(w, request, &data)

		if err == nil {
			t.Error("ReadJSON() expected an error, got nil")
		} else if err.Error() != "http: request body too large" {
			t.Errorf("ReadJSON() error message = %s, want 'http: request body too large'", err.Error())
		}
	})

	t.Run("Empty Request Body", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{}

		jsonData := `{}`

		request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		err = ReadJSON(w, request, &data)

		if err != nil {
			t.Errorf("ReadJSON() error = %v, want nil", err)
		}

		if data.Name != "" {
			t.Errorf("ReadJSON() data.Name = %s, want an empty string", data.Name)
		}
	})
}

func TestWriteJSON(t *testing.T) {
	t.Parallel()

	t.Run("Valid JSON", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{Name: "John"}

		w := httptest.NewRecorder()

		err := WriteJSON(w, http.StatusOK, data)

		if err != nil {
			t.Errorf("WriteJSON() error = %v, want nil", err)
		}

		if w.Code != http.StatusOK {
			t.Errorf("WriteJSON() status code = %d, want %d", w.Code, http.StatusOK)
		}

		// Parse the expected and actual JSON responses
		var expectedJSON, actualJSON map[string]interface{}
		if err := json.Unmarshal([]byte(`{
			"name": "John"
		}`), &expectedJSON); err != nil {
			t.Fatalf("Failed to unmarshal expected JSON: %v", err)
		}

		if err := json.Unmarshal(w.Body.Bytes(), &actualJSON); err != nil {
			t.Fatalf("Failed to unmarshal actual JSON: %v", err)
		}

		// Compare the parsed JSON content
		if !reflect.DeepEqual(expectedJSON, actualJSON) {
			t.Errorf("WriteJSON() response JSON = %+v, want %+v", actualJSON, expectedJSON)
		}
	})

	t.Run("Invalid JSON Encoding", func(t *testing.T) {
		invalidData := make(chan int) // Unsupported type for JSON encoding

		w := httptest.NewRecorder()

		err := WriteJSON(w, http.StatusOK, invalidData)

		if err == nil {
			t.Error("WriteJSON() expected an error, got nil")
		}
	})

	t.Run("Additional Headers", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
		}{Name: "John"}

		w := httptest.NewRecorder()

		headers := make(http.Header)
		headers.Add("Custom-Header", "CustomValue")

		err := WriteJSON(w, http.StatusOK, data, headers)

		if err != nil {
			t.Errorf("WriteJSON() error = %v, want nil", err)
		}

		if w.Header().Get("Custom-Header") != "CustomValue" {
			t.Errorf("WriteJSON() custom header not set correctly")
		}
	})
}
