package handlers

import (
	"net/http"
	"strconv"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/utils"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// update a specific product
// responses:
// 	200: productResponse
//  400: errorResponse
//  404: errorResponse
//  500: errorResponse
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.JSONError(rw, utils.BadRequestError("something went wrong"))
		return
	}
	
	p.l.Println("PUT")

	rw.Header().Add("Content-Type","application/json")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		utils.JSONError(rw, utils.NotFoundError("product"))
		return
	}
	if err != nil {
		utils.JSONError(rw, utils.InternalServerError("something went wrong"))
		return
	}

	err = prod.ToJSON(rw)
	if err != nil {
		utils.JSONError(rw, utils.InternalServerError("something went wrong"))
	}
}