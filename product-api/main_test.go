package main

import (
	"fmt"
	"github.com/piyush1146115/go-microservice/product-api/sdk/client"
	"github.com/piyush1146115/go-microservice/product-api/sdk/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}
