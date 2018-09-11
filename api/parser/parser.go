package parser

import (
	"fmt"
	"net/http"

	"github.com/liftedkilt/foothillnorthchur.ch/api/db"

	"golang.org/x/net/html"
)

// Parse takes a url as a string and returns a slice of struct Video
func Parse(url string) (playlist []db.Video) {
	playlist = scrapePage(url)

	return
}

// Helper function to pull the titledata attribute from a Token
func getTagAttr(t html.Token) (ok bool, video db.Video) {
	// Iterate over all of the Token's attributes until we find an "data-title"
	for _, a := range t.Attr {
		if a.Key == "data-video-id" {
			video.VideoID = a.Val
			fmt.Println(a)
			ok = true
		} else if a.Key == "data-title" {
			video.VideoTitle = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, dataTitle) as defined in
	// the function definition
	return
}

// Extract all titles from a youtube playlist
func scrapePage(url string) (videos []db.Video) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is a <tr> tag
			isTableRow := t.Data == "tr"
			if !isTableRow {
				continue
			}

			// Extract the data-title value, if there is one
			// ok, videoDataTitle := getHref(playlistID, t)
			ok, video := getTagAttr(t)
			if !ok {
				continue
			}

			videos = append(videos, video)
		}
	}
}
