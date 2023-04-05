package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//開発環境
	stgServer()
}
func stgServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	log.Printf("port : %s",port)
	mux.HandleFunc("/hello", HelloWorldHandler)


	if err := http.ListenAndServe(fmt.Sprintf(":%s",port), mux); err != nil {
		log.Fatal(err)
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("started server")
	log.Printf(r.RemoteAddr)
	fmt.Fprintf(w,"Hello %s!",r.URL.Path[1:])
}