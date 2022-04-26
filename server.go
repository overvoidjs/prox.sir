package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Перенаправляем запросы куда надо
func reversProxy(w http.ResponseWriter, r *http.Request) {

	host := r.Host
	knowThisServer := os.Getenv(host)

	if knowThisServer == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404")
	} else {

		url, err := url.Parse(knowThisServer)
		if err != nil {
			log.Println(err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	}

}

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error load .env file")
	}

	http.HandleFunc("/", reversProxy)
	fmt.Println("Prox.Sir: Ready")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Prox.Sir: ", err)
	}
}
