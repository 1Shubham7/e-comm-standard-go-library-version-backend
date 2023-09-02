package handlers

import (
	"log"
	"net/http"
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


}