package parser

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Video is a struct that contains the video order as an int and the videoID as a string
type Video struct {
	Order   int    `json:"order"`
	VideoID string `json:"videoID"`
}

// Parse takes a url as a string and returns a slice of struct Video
func Parse(playlistID string) (playlist []Video) {
	videoIDs := scrapePage(playlistID)

	videoIDs = sliceUniqMap(videoIDs)

	playlist = []Video{}

	for i, videoID := range videoIDs {
		playlist = append(playlist, Video{Order: i + 1, VideoID: videoID})
	}

	return
}

// Helper function to pull the titledata attribute from a Token
func getDataTitle(t html.Token) (ok bool, dataTitle string) {
	// Iterate over all of the Token's attributes until we find an "data-title"
	for _, a := range t.Attr {
		if a.Key == "data-title" {
			dataTitle = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, dataTitle) as defined in
	// the function definition
	return
}

// remove duplicate IDs
func sliceUniqMap(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

// Helper function to pull the titledata attribute from a Token
func getHref(playlistID string, t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "data-title"
	for _, a := range t.Attr {
		if a.Key == "href" {
			if strings.HasPrefix(a.Val, "/watch?v=") && strings.Contains(a.Val, playlistID) {
				// href = a.Val
				href = strings.TrimPrefix(a.Val, "/watch?v=")
				href = strings.Split(href, "&")[0]
				ok = true
			}
		}
	}

	// "bare" return will return the variables (ok, dataTitle) as defined in
	// the function definition
	return
}

// Extract all titles from a youtube playlist
func scrapePage(playlistID string) (titles []string) {

	url := "https://youtube.com/playlist?list=" + playlistID

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

			// Check if the token is a <a> tag
			isTableRow := t.Data == "a"
			if !isTableRow {
				continue
			}

			// Extract the data-title value, if there is one
			ok, title := getHref(playlistID, t)
			if !ok {
				continue
			}

			titles = append(titles, title)
		}
	}
}
