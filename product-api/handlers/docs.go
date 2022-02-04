// Package classification Product API
//
// Documentation for Product API
//
// 	Schemes: http
//	BasePath: /
// 	Version: 0.0.1
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// 	- application/text
//
// swagger:meta
package handlers

import (
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/utils"
)

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct{
	// All products in the system
	// in: body
	Body []data.Product
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct{
	// Single product in the system
	// in: body
	Body data.Product
}

// custom json error
// swagger:response errorResponse
type errorResponseWrapper struct{
	// Description of error(s)
	// in: body
	Body utils.CustomError
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}