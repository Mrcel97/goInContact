package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

type Link struct {
	url string
}

var topbar []Link

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://golang.org/", 301)
}

func getLinks(attr []html.Attribute) string {
	for _, a := range attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func parseLinks(respBody io.ReadCloser) []Link {
	var links []Link

	z := html.NewTokenizer(respBody)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				links = append(links, Link{getLinks(t.Attr)})
			}
		}
	}
}

func formatTopbar(links []Link) []Link {
	for link := range links {
		topbar = append(topbar, Link{"https://golang.org" + links[link].url})
	}
	return topbar
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Golang Website Topbar Links: </h1>")
	for link := range topbar {
		fmt.Fprintf(w, "<a href="+topbar[link].url+">"+topbar[link].url+"</a>")
		fmt.Fprint(w, "<br>")
	}
}

func main() {
	resp, _ := http.Get("https://golang.org/") // We can use '_' if we don't want a given parameter. Is an excelent way to avoid "Variable not used" errors

	var links = parseLinks(resp.Body)
	links = formatTopbar(links[3:8])

	resp.Body.Close()

	http.Handle("/", http.FileServer(http.Dir("./app")))
	http.HandleFunc("/info", redirect)
	http.HandleFunc("/links", linksHandler)
	http.ListenAndServe(":8000", nil)
}

/*
Array vs Slice:
	[5]type == array
	[]type == slice
*/
