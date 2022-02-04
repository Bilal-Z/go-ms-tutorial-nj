package handlers

import (
	"net/http"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/utils"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse
//  500: errorResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("GET")

	rw.Header().Add("Content-Type","application/json")

	// list of products
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		utils.JSONError(rw, utils.InternalServerError("something went wrong"))
		return
	}
}