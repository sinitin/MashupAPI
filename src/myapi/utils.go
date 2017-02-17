package main

import "net/http"

type Summary struct {
	Mbid        string   `json:"mbid"`
	Description string   `json:"description"`
	Albums      []*Album `json:"albums"`
}

type Album struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Image string `json:"image"`
}

type CoverArt struct {
	Images []struct {
		Approved   bool   `json:"approved"`
		Back       bool   `json:"back"`
		Comment    string `json:"comment"`
		Edit       int    `json:"edit"`
		Front      bool   `json:"front"`
		ID         string `json:"id"`
		Image      string `json:"image"`
		Thumbnails struct {
			Large string `json:"large"`
			Small string `json:"small"`
		} `json:"thumbnails"`
		Types []string `json:"types"`
	} `json:"images"`
	Release string `json:"release"`
}

type Relation struct {
	Ended           bool            `json:"ended"`
	TargetType      string          `json:"target-type"`
	Direction       string          `json:"direction"`
	Attributes      []interface{}   `json:"attributes"`
	URL             URL             `json:"url"`
	AttributeValues AttributeValues `json:"attribute-values"`
	End             interface{}     `json:"end"`
	Begin           interface{}     `json:"begin"`
	Type            string          `json:"type"`
	TypeID          string          `json:"type-id"`
	TargetCredit    string          `json:"target-credit"`
	SourceCredit    string          `json:"source-credit"`
}

type URL struct {
	Resource string `json:"resource"`
	ID       string `json:"id"`
}

type AttributeValues struct {
}

type ReleaseGroup struct {
	SecondaryTypeIds []string `json:"secondary-type-ids"`
	FirstReleaseDate string   `json:"first-release-date"`
	PrimaryType      string   `json:"primary-type"`
	PrimaryTypeID    string   `json:"primary-type-id"`
	SecondaryTypes   []string `json:"secondary-types"`
	Disambiguation   string   `json:"disambiguation"`
	ID               string   `json:"id"`
	Title            string   `json:"title"`
}

type WikiIntro struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Normalized []struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"normalized"`
		Pages map[string]interface{} `json:"pages"`
	} `json:"query"`
}

type AlbumInput struct {
	url   string
	title string
	id    string
}

type AlbumResponse struct {
	response *http.Response
	err      error
	title    string
	id       string
	coverArt CoverArt
}

type ArtistInfo struct {
	Area struct {
		Disambiguation   string   `json:"disambiguation"`
		ID               string   `json:"id"`
		Iso_3166_1_codes []string `json:"iso-3166-1-codes"`
		Name             string   `json:"name"`
		Sort_name        string   `json:"sort-name"`
	} `json:"area"`
	BeginArea struct {
		Disambiguation string `json:"disambiguation"`
		ID             string `json:"id"`
		Name           string `json:"name"`
		Sort_name      string `json:"sort-name"`
	} `json:"begin_area"`
	Country        string        `json:"country"`
	Disambiguation string        `json:"disambiguation"`
	EndArea        interface{}   `json:"end_area"`
	Gender         interface{}   `json:"gender"`
	Gender_id      interface{}   `json:"gender-id"`
	ID             string        `json:"id"`
	Ipis           []interface{} `json:"ipis"`
	Isnis          []string      `json:"isnis"`
	Life_span      struct {
		Begin string `json:"begin"`
		End   string `json:"end"`
		Ended bool   `json:"ended"`
	} `json:"life-span"`
	Name           string         `json:"name"`
	Relations      []Relation     `json:"relations"`
	Release_groups []ReleaseGroup `json:"release-groups"`
	Sort_name      string         `json:"sort-name"`
	Type           string         `json:"type"`
	Type_id        string         `json:"type-id"`
}
