package handlers

import (
	"context"
	"net/http"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/data"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/utils"
	"github.com/go-playground/validator/v10"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request){
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			utils.JSONError(rw, utils.InternalServerError("error reading product"))
			return
		}
		// validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			utils.JSONError(rw, utils.RequestValidationError(
				err.(validator.ValidationErrors),
				))
			return
		}

		ctx := context.WithValue(r.Context(),KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}