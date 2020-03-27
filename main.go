package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/reujab/wallpaper"
)

// BingURL represents URL to bing website.
const BingURL = "http://www.bing.com"

func main() {
	res, err := http.Get(BingURL)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	compiled, err := regexp.Compile("id=\"downloadLink\".*?href=\"(.*?)\"")
	if err != nil {
		log.Fatal(err)
	}

	matched := compiled.FindStringSubmatch(bodyString)
	pictureURL := matched[len(matched)-1]

	builtURL := strings.Builder{}
	builtURL.WriteString(BingURL)
	builtURL.WriteString(pictureURL)

	err = wallpaper.SetFromURL(builtURL.String())
	if err != nil {
		log.Fatal(err)
	}
}
