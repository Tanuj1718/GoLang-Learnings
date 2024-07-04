package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct{
	ID string `json: "id"`
	OriginalURL string `json: "original_url"`
	ShortURL string `json: "short_url"`
	CreationDate time.Time `json: "creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string  {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) //convert the original url string to a byte slice

	data := hasher.Sum(nil) //converted into array or slice

	hash := hex.EncodeToString(data) // converted into string
	fmt.Println("Encoded: ", hash)

	//converted into 8 size
	fmt.Println("Final string:", hash[:8])
	return hash[:8]
}

func createURL(originalURL string) string  {
	shortURL := generateShortURL(originalURL)
	id := shortURL // using short url as the id
	//saving the short url in database/here map
	urlDB[id] = URL{
		ID: id,
		OriginalURL: originalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error)  {
	url , ok := urlDB[id]
	if(!ok){
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func shortURLHandler(w http.ResponseWriter, r *http.Request){
	var data struct {
		URL string `json: "url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if(err!=nil){
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortUrl := createURL(data.URL)
	//sending the response or printing in postman
	response := struct{
		ShortURL string `json:"short_url"`
	}{ShortURL: shortUrl}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if(err!=nil){
		http.Error(w, "Invalid request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	//fmt.Println("Starting URL shortener...")
	//generateShortURL("https://github.com/prince")


	//register the handler function to handle all requests to the root URL ("/")
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)


	//starting the http server
	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if(err!=nil){
		fmt.Println("Error on starting server:", err)
	}
}