package main

import (
	"fmt"
)

func getArtistInfo(mbid string) (ArtistInfo, error) {
	var artistInfo ArtistInfo
	err := getJson(fmt.Sprintf("http://musicbrainz.org/ws/2/artist/%s%s", mbid, "?&fmt=json&inc=url-rels+release-groups"), &artistInfo)
	if err != nil {
		return artistInfo, err
	}

	return artistInfo, nil
}

func getCoverArt(releaseGroups []ReleaseGroup) ([]*Album, error) {
	//pick up the cover art, the album title and id
	albumInputList := []AlbumInput{}
	for _, releaseGroups := range releaseGroups {
		albumInput := new(AlbumInput)
		albumInput.url = fmt.Sprintf("http://coverartarchive.org/release-group/%s", releaseGroups.ID)
		albumInput.id = releaseGroups.ID
		albumInput.title = releaseGroups.Title
		albumInputList = append(albumInputList, *albumInput)
	}

	responses, _ := asyncHttpGets(albumInputList)
	// TODO: decide if failure to fetch coer art should return a 404
	/*if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}*/
	albums := []*Album{}

	for _, r := range responses {
		album := new(Album)
		for _, image := range r.coverArt.Images {
			if image.Front == true {
				album.Image = image.Image
				break
			}
		}
		album.ID = r.id
		album.Title = r.title
		albums = append(albums, album)
	}

	return albums, nil
}
