package routes

import (
	"github.com/gorilla/mux"
	"script_pilot/internal/api"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/capcut_input", api.GenerateScriptHandler).Methods("POST")
	return router
}
