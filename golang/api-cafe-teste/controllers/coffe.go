package controllers

import (
	"api-cafe/helpers"
	"api-cafe/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var models services.Models
var coffe = models.Coffe

func GetAllCoffes(w http.ResponseWriter, r *http.Request) {
	var coffees services.Coffe
	all, err := coffees.GetAllCoffess()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffe, err := coffe.GetCoffeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w, http.StatusOK, coffe)
}

func CreateCoffe(w http.ResponseWriter, r *http.Request) {
	var coffeeResp services.Coffe
	err := json.NewDecoder(r.Body).Decode(&coffeeResp)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w, http.StatusOK, coffeeResp)
	coffeeCreated, err := coffe.CreateCoffe(coffeeResp)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusOK, coffeeCreated)
	}
}

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var coffe services.Coffe
	err := json.NewDecoder(r.Body).Decode(&coffe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJson(w, http.StatusOK, coffe)
	coffeObj, err := coffe.UpdateCoffe(id, coffe)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusOK, coffeObj)
	}
}

func DeleteCoffe(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffe.DeleteCoffe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	helpers.WriteJson(w, http.StatusOK, "succesfull deleteion")
}
