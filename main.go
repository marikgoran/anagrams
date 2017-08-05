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
	a.Anagrams = getPerms(word)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(a)
}

func getPerms(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1] // current char
	remStr := str[1:] // remaining string

	perms := getPerms(remStr) // get perms for remaining string

	allPerms := make([]string, 0) // array to hold all perms of the string based on perms of substring

	// for every perm in the perms of substring
	for _, perm := range perms {
	        // add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}

	return allPerms
}

// Insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}
