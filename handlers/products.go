package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/1shubham7/e-comm/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

// anything adds to the handler interface  needs to have a serve http method
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
// we will now use encoding/json package to convert the struct we ahve into a json representative
	listOfProducts := data.GetProducts()
	// now listOfProducts contains the list of products, but how do we return that to the user
	// we do that by converting listOfProducts into json. and this the the way

	d, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "unable to convert data to json", http.StatusInternalServerError)
	}
	rw.Write(d)
}