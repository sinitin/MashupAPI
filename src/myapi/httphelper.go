package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
)

func getJson(url string, target interface{}) error {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "myapi/3.0 ( sini.tinfors@gmail.com )")

	r, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response code, got: %d expected: 200, for url %s", r.StatusCode, url)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func asyncHttpGets(input []AlbumInput) ([]*AlbumResponse, error) {

	ch := make(chan *AlbumResponse, len(input))
	responses := []*AlbumResponse{}
	for _, i := range input {
		url := i.url
		id := i.id
		title := i.title
		go func(url string, id string, title string) {
			resp, err := http.Get(url)
			coverArt := new(CoverArt)
			json.NewDecoder(resp.Body).Decode(coverArt)
			resp.Body.Close()
			ch <- &AlbumResponse{resp, err, title, id, *coverArt}
		}(url, id, title)
	}

	timeout := time.After(10000 * time.Millisecond)

	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == len(input) {
				return responses, nil
			}
		case <-timeout:
			return responses, fmt.Errorf("Timeout waiting for cover art")
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
