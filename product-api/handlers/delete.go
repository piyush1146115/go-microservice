package handlers

import (
	"github.com/piyush1146115/go-microservice/product-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 201: noContentResponse
// 404: errorResponse
// 501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)
		rw.WriteHeader(http.StatusInternalServerError)

		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}
