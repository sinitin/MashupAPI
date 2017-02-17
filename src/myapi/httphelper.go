package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getJson(url string, target interface{}) (error) {

	//fmt.Printf("Hello, the URL I go was: %s", url)

	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected http response code, got: %d expected: 200", r.StatusCode)
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

	/*for {
		select {
		case r := <-ch:
			fmt.Println("Got at cover!!!")
			responses = append(responses, r)
			if len(responses) == len(input) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}

	return responses*/

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
