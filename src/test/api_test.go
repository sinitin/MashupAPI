package test

import (
	"encoding/json"
	"fmt"
	"myapi"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetArtistAndMbid(t *testing.T) {

	testcases := make(map[string]int)

	testcases["5b11f4ce-a62d-471e-81fc-a69a8278c7da"] = 200 //Ok request
	testcases["2cd02888-5f57-4973-892e-53937a554d09"] = 404 //Random UUID
	testcases["malformeduuid"] = 400                        //Malformed UUID

	var myClient = &http.Client{Timeout: 120 * time.Second}

	for mbid, expectedresp := range testcases {

		r, err := myClient.Get(fmt.Sprintf("http://localhost:8080/musicinfo/%s", mbid))
		if err != nil {
			t.Errorf("Something went wrong: %v", err)
		}

		if r.StatusCode != expectedresp {
			fmt.Println(mbid)
			t.Errorf("Wrong http response code, expected: %d, got: %d, mbid %s", expectedresp, r.StatusCode, mbid)
		}

		if r.StatusCode == 200 {
			defer r.Body.Close()

			var summary main.Summary
			json.NewDecoder(r.Body).Decode(&summary)

			if summary.Mbid != mbid {
				t.Errorf("Wrong mbid in return, got: %s", summary.Mbid)
			}
		}
	}

}

func TestGetBritneySpears(t *testing.T) {

	testcases := make(map[string]int)

	testcases["45a663b5-b1cb-4a91-bff6-2bef7bbfdd76"] = 200 //Britney Spears

	var myClient = &http.Client{Timeout: 120 * time.Second}

	for mbid, expectedresp := range testcases {

		r, err := myClient.Get(fmt.Sprintf("http://localhost:8080/musicinfo/%s", mbid))
		if err != nil {
			t.Errorf("Something went wrong: %v", err)
		}

		if r.StatusCode != expectedresp {
			t.Errorf("Wrong http response code, expected: %d", expectedresp)
		}

		defer r.Body.Close()

		var summary main.Summary
		json.NewDecoder(r.Body).Decode(&summary)

		if strings.Compare(summary.Mbid, mbid) != 0 {
			t.Errorf("Wrong mbid in return, expected: %s, got: %s", mbid, summary.Mbid)
		}

		if len(summary.Description) == 0 {
			t.Errorf("Failed to fetch description")
		}

		found := false
		for _, album := range summary.Albums {
			if album.Title == "Circus" && album.Image == "http://coverartarchive.org/release/a40a2430-2483-4936-9293-1a5b4d6b1b22/12099364119.jpg" {
				found = true
				break
			}
		}

		if found == false {
			t.Errorf("Failed to fetch the cover art")
		}
	}
}
