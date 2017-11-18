package main

import (
	"encoding/json"
	// "fmt"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Items struct {
	Items []Result `json:"items"`
}

type Result struct {
	Title string `json:"title"`
	Arg   string `json:"arg"`
}

func main() {
	passDir := os.Getenv("PASSWORD_STORE_DIR")

	if len(passDir) == 0 {
		passDir = os.Getenv("HOME") + "/.password-store/"
	}

	query := strings.Replace(os.Args[1], " ", "/", -1)

	files, err := filepath.Glob(passDir + "**/*.gpg")

	if err != nil {
		log.Fatal(err)
	}

	matches := fuzzy.Find(query, files)
	results := make([]Result, len(matches))

	r := strings.NewReplacer(passDir, "", ".gpg", "")
	for i, match := range matches {
		title := r.Replace(match)
		results[i] = Result{title, match}
	}

	items := Items{results}

	b, err := json.Marshal(items)

	os.Stdout.Write(b)
}
