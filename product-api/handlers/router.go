package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRouter(r *mux.Router, l *log.Logger){
	ph := NewProduct(l)

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", ph.GetProducts)
	getR.HandleFunc("", ph.GetProducts)

	postR := r.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/", ph.AddProduct)
	postR.HandleFunc("", ph.AddProduct)
	postR.Use(ph.MiddlewareProductValidation)

	putR := r.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putR.Use(ph.MiddlewareProductValidation)
}