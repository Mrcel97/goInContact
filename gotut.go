package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://golang.org/", 301)
}

func main() {
	resp, _ := http.Get("https://golang.org/") // format: response, error := http.Get(URL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()

	http.Handle("/", http.FileServer(http.Dir("./app")))
	http.HandleFunc("/info", redirect)
	http.ListenAndServe(":8000", nil)
}

/*
Line 8: Underscore '_' can be used if we will not use of a received parameter.
*/
