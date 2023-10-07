package helpers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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

		expectedJSON := `{
			"name": "John"
		}`

		if w.Body.String() != expectedJSON {
			t.Errorf("WriteJSON() response body = %s, want %s", w.Body.String(), expectedJSON)
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
}
