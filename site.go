package main

import (
	"github.com/golang/groupcache/lru"
	"io/ioutil"
	"log"
	"net/http"
)

type SiteCach struct {
	url   string
	cache lru.Cache
}

func (site *SiteCach) ImageCashing() {
	err := site.caching(site.url+"/image.jpg", &site.cache)
	if err != nil {
		log.Fatal(err)
	}
}

func (site *SiteCach) HtmlCashing() {
	err := site.caching(site.url, &site.cache)
	if err != nil {
		log.Fatal(err)
	}
}

func (site *SiteCach) caching(url string, cache *lru.Cache) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	cache.Add(url, byteBody)

	return nil
}
