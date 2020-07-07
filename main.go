package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/reujab/wallpaper"
)

// BingURL represents URL to bing website.
const BingURL = "http://www.bing.com"

// BingPhotoOfTheDayURL represents URL to Bing photo of the day API endpoint.
const BingPhotoOfTheDayURL = "/HPImageArchive.aspx?format=js&n=1"

// BingPhotoOfTheDayResponse represents a struct with Bing photo of the day endpoint response data.
type BingPhotoOfTheDayResponse struct {
	Images []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		URL           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Title         string        `json:"title"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

func main() {
	sb := strings.Builder{}
	sb.WriteString(BingURL)
	sb.WriteString(BingPhotoOfTheDayURL)

	res, err := http.Get(sb.String())
	if err != nil {
		log.Fatal(err)
	}

	bingPhotoResponse := &BingPhotoOfTheDayResponse{}
	err = json.NewDecoder(res.Body).Decode(bingPhotoResponse)
	if err != nil {
		log.Fatal(err)
	}

	if len(bingPhotoResponse.Images) < 1 {
		log.Fatalf("Images missing in Bing API response!")
	}

	sb = strings.Builder{}
	sb.WriteString(BingURL)
	sb.WriteString(bingPhotoResponse.Images[0].URL)

	err = wallpaper.SetFromURL(sb.String())
	if err != nil {
		log.Fatal(err)
	}
}
