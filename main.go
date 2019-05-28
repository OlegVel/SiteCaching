package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url = "http://108.61.245.170"

type SiteCash struct {
	html  bytes.Buffer
	image bytes.Buffer
}

func main() {
	var site = &SiteCash{}
	site.HtmlCashing()
	site.ImageCashing()
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for t := range ticker.C {
			site.HtmlCashing()
			site.ImageCashing()
			fmt.Println("Tick at", t)
		}
	}()

	http.Handle("/", SiteHandler())
	http.Handle("/image.jpg", ImageHandler(site))
	if err := http.ListenAndServe(":1080", nil); err != nil {
		log.Fatal(err)
	}
}

func SiteHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		bBytes, err := ioutil.ReadAll(request.Body)
		ss := string(bBytes)
		fmt.Println("Header: ", request.Header)
		fmt.Println("Body: ", ss)

		client := &http.Client{}
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		writer.Write(bodyBytes)
		resp.Body.Close()

	}
}

func ImageHandler(site *SiteCash) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var temp = bytes.Buffer{}
		tee := io.TeeReader(&site.image, writer)
		_, err := io.Copy(&temp, tee)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(&site.image, &temp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (site *SiteCash) ImageCashing() {
	client := &http.Client{}
	resp, err := client.Get(url + "/image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(&site.image, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (site *SiteCash) HtmlCashing() {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(&site.html, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
