package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	var site = &SiteCach{
		url: "http://108.61.245.170",
	}
	site.HtmlCashing()
	site.ImageCashing()

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for t := range ticker.C {
			site.HtmlCashing()
			site.ImageCashing()
			fmt.Println("Updated at", t)
		}
	}()

	http.Handle("/", SiteHandler(site))
	http.Handle("/image.jpg", ImageHandler(site))
	if err := http.ListenAndServe(":1080", nil); err != nil {
		log.Fatal(err)
	}
}

func SiteHandler(site *SiteCach) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(site.cache.Len())
		resp, ok := site.cache.Get(site.url)
		if !ok {
			log.Fatal("error getting")
		}
		response := bytes.NewReader(resp.([]byte))

		_, err := io.Copy(writer, response)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func ImageHandler(site *SiteCach) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		resp, ok := site.cache.Get(site.url + "/image.jpg")
		if !ok {
			log.Fatal("error getting")
		}
		response := bytes.NewReader(resp.([]byte))

		_, err := io.Copy(writer, response)
		if err != nil {
			log.Fatal(err)
		}
	}
}
