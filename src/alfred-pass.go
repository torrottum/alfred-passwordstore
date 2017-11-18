package main

import (
	"encoding/json"
	"github.com/sahilm/fuzzy"
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
	if len(os.Args) == 1 {
		// just exit silently if no argument
		os.Exit(0)
	}

	query := strings.Replace(os.Args[1], " ", "/", -1)

	passDir := os.Getenv("PASSWORD_STORE_DIR")
	if len(passDir) == 0 {
		passDir = os.Getenv("HOME") + "/.password-store/"
	}

	files, err := filepath.Glob(passDir + "**/*.gpg")
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReplacer(passDir, "", ".gpg", "")
	for i, file := range files {
		files[i] = r.Replace(file)
	}

	matches := fuzzy.Find(query, files)
	results := make([]Result, len(matches))
	for i, match := range matches {
		title := match.Str
		results[i] = Result{title, passDir + title + ".gpg"}
	}

	b, err := json.Marshal(Items{results})
	os.Stdout.Write(b)
}
