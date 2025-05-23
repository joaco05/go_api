package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go_api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting API...")

	error := http.ListenAndServe("localhost:8000", r)
	if error != nil {
		log.Error(error)
	}
}