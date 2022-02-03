package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("GET")

	// list of products
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "something went wrong", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("POST")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

	rw.WriteHeader(http.StatusCreated)
	err := prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusInternalServerError)
	}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusBadRequest)
		return
	}
	
	p.l.Println("PUT")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusInternalServerError)
		return
	}

	err = prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request){
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "error reading product", http.StatusInternalServerError)
			return
		}
		// validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("error validating product %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(),KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}