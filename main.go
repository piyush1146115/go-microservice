package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello world")

		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(writer, "Hello %s\n", data)
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
