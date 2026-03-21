package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request){
		// 
	})

	http.HandleFunc("/shorten", func(writer http.ResponseWriter, req *http.Request){
		// Get the url to shorten from the request
		url := req.FormValue("url")
		fmt.Println("Payload: ", url)
		// Shorten the url
		shortURL := utils.GetShortCode()
		fullShortURL := fmt.Sprintf("http://localhost:8080/r/%s", shortURL)
		// Generated short URL
		fmt.Printf("Generated short URL: %s\n", shortURL)
	})
}