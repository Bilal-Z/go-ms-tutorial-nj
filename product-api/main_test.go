package main

import (
	"fmt"
	"testing"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/sdk/client"
	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/sdk/client/products"
)

func TestClient(t *testing.T){
	cfg := client.DefaultTransportConfig().WithHost("localhost:5000")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prod)
}