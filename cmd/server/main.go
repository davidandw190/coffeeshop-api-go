package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	// Model 1
	// ...
	// Model n
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	log.Println("Server: API listening on port", port)

	s := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO: add router
	}

	return s.ListenAndServe()

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	c := Config{
		Port: os.Getenv("PORT"),
	}

	// TODO: connenction to db

	app := &Application{
		Config: c,
		// TODO: add the models
	}

	if err = app.Serve(); err != nil {
		log.Fatal(err)
	}
}
