package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request){
		fmt.Fprintln(writer, "Welcome to Chibi URL!")
	})

	fmt.Println("Server is running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}