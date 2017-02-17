package main

import (
	"fmt"
	"log"
	"strings"
)

func getWikiDescription(relations []Relation) (string, error) {

	//get the wikipedia name of the artist or band from the musicbrainz response
	var url string
	for _, rel := range relations {
		if strings.Compare(rel.Type, "wikipedia") == 0 {
			url = rel.URL.Resource
			break
		}
	}

	if len(url) == 0 {
		return "", fmt.Errorf("No url to wikipedia info found")
	}

	parts := strings.SplitAfter(url, "/")
	l := len(parts)
	artistname := parts[l-1]

	//pick up the wikipedia introduction text
	wikiIntro := new(WikiIntro)
	err := getJson(fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s", artistname), wikiIntro)
	if err != nil {
		log.Print(err)
		return "", err
	}

	//TODO: fix this not so pretty solution for handlning a dynamic pages field in the json response from wikipedia
	var m interface{} = wikiIntro.Query.Pages
	wikimap := m.(map[string]interface{})
	for _, value := range wikimap {
		var unknown interface{} = value
		nestedwikimap := unknown.(map[string]interface{})
		for _, value = range nestedwikimap {
			if val, ok := nestedwikimap["extract"]; ok {
				var extract interface{} = val
				description := extract.(string)
				return description, nil
			}
		}
	}

	return "", fmt.Errorf("Unable to fetch description of artist")
}
