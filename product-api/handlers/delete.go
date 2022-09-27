package handlers

import (
	"github.com/gorilla/mux"
	"github.com/piyush1146115/go-microservice/product-api/data"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	// this will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}
