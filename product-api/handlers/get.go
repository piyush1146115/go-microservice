package handlers

import (
	"github.com/piyush1146115/go-microservice/product-api/data"
	"net/http"
)

// swagger:route GET /products products listProducts

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}
}
