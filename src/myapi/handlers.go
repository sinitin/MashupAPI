package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func MusicInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	mbid := vars["MBID"]
	var summary Summary

	if ok, err := checkOkMbid(mbid); !ok || (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check if we already have the cashed data
	if err := RepoCheckMbid(mbid); err == nil {
		summary = RepoGetSummary(mbid)

		//fmt.Println("     Yaaaay found data in cache    ")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(summary); err != nil {
			panic(err)
		}

		return

	}

	//check if the artist exists and pick up the albums
	artistInfo, err := getArtistInfo(mbid)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	summary.Mbid = mbid

	//pick up the description from wikipedia
	var relations []Relation
	relations = artistInfo.Relations
	summary.Description, _ = getWikiDescription(relations)
	// TODO: decide if an empty description from wikipedia should return a 404
	/*if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

	var releaseGroups []ReleaseGroup
	releaseGroups = artistInfo.Release_groups
	summary.Albums, _ = getCoverArt(releaseGroups)
	// TODO: decide if no cover art should return a 404
	/*if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
	}*/

	RepoAddSummary(summary)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		panic(err)
	}
}
