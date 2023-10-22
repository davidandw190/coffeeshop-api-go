package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/davidandw190/coffeeshop-api-go/helpers"
	"github.com/davidandw190/coffeeshop-api-go/services"
)

var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	coffees, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": coffees})
}

// POST/coffees/coffee
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}
