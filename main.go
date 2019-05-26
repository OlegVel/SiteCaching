package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", SiteHandler())
	fmt.Println("Hello!")
	if err := http.ListenAndServe(":1080", nil); err != nil {
		log.Fatal(err)
	}
}

func SiteHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Body: ", request.Body)
	}
}
