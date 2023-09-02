package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("I love GeeksforGeeks")

		d, err := ioutil.ReadAll(r.Body)

		if err!= nil{
			http.Error(rw, "Sorry, Try again later", http.StatusBadRequest)
			return
		}
		log.Printf("Data : %s\n", d)
		fmt.Fprintf(rw, "\nHello There %s", d)
	})
	http.ListenAndServe(":6000", nil)
}
