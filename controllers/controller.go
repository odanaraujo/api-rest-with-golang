package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-rest-with-golang/models"
	"net/http"
	"strconv"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Home page")
}

func AllPersonalitys(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(models.Personalitys)
}

func FindPersonalityById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	for _, personality := range models.Personalitys {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(writer).Encode(personality)
		}
	}
}
