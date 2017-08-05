package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

type AnagramResult struct {
    Word	   string
    Anagrams []string
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/anagrams/", anagrams )

    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func anagrams (w http.ResponseWriter, r *http.Request) {
	word := strings.SplitN(r.URL.Path, "/", 3)[2]
	var a AnagramResult

	if  len(word)>5 {
		http.Error(w, "Currently not supporting anagrams with more then 5 letters", http.StatusInternalServerError)
		return
	}

	a.Word = word
	a.Anagrams = []string{"A","B",word}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(a)
}

