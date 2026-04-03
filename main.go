package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kk-ami/Chibi-URL/utils"
)

var ctx = context.Background()

func main() {
	dbClient := utils.NewRedisClient()
	if dbClient == nil {
		fmt.Println("Failed to connect to redis")
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(writer, nil)
	})

	http.HandleFunc("/shorten", func(writer http.ResponseWriter, req *http.Request) {
		// Get the url to shorten from the request
		url := req.FormValue("url")
		fmt.Println("Payload: ", url)
		// Shorten the url
		shortURL := utils.GetShortCode()
		fullShortURL := fmt.Sprintf("http://localhost:8080/r/%s", shortURL)
		// Generated short URL
		fmt.Printf("Generated short URL: %s\n", shortURL)

		utils.SetKey(&ctx, dbClient, shortURL, url, 0)
		fmt.Fprintf(writer,
			`<p class="mt-4 text-green-600">Shortened URL: <a href="/r/%s" class="underline">%s</a></p>`, shortURL, fullShortURL)
	})

	http.HandleFunc("/r/{code}", func(writer http.ResponseWriter, req *http.Request) {
		key := req.PathValue("code")
		if key == "" {
			http.Error(writer, "Invalid URL", http.StatusBadRequest)
			return
		}
		longURL, err := utils.GetLongURL(&ctx, dbClient, key)
		if err != nil {
			http.Error(writer, "shortned URL not found", http.StatusNotFound)
			return
		}
		http.Redirect(writer, req, longURL, http.StatusPermanentRedirect)

	})
	http.ListenAndServe(":8080", nil)

}
