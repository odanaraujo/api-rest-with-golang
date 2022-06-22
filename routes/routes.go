package routes

import (
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-rest-with-golang/controllers"
	"log"
	"net/http"
)

func LoadRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalitys).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
