package parser

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Song is a struct that contains the song order as an int and a title as a string
type Song struct {
	Order int    `json:"order"`
	Title string `json:"title"`
}

// Parse takes a url as a string and returns a slice of struct Playlist
func Parse(url string) (playlist []Song) {
	titles := scrapePage(url)

	playlist = []Song{}

	for i, title := range titles {
		playlist = append(playlist, Song{Order: i + 1, Title: title})
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

// Extract all titles from a youtube playlist
func scrapePage(url string) (titles []string) {
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
			ok, title := getDataTitle(t)
			if !ok {
				continue
			}

			titles = append(titles, title)
		}
	}
}
