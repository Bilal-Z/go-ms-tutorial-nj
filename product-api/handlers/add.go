package handlers

import (
	"net/http"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/utils"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	201: productResponse
//  400: errorResponse
//  500: errorResponse
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("POST")

	rw.Header().Add("Content-Type","application/json")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

	rw.WriteHeader(http.StatusCreated)
	err := prod.ToJSON(rw)
	if err != nil {
		utils.JSONError(rw, utils.InternalServerError("something went wrong"))
	}
}