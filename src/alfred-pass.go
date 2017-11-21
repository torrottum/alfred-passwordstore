package main

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sahilm/fuzzy"
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

	files := []string{}
	err := filepath.Walk(passDir, func(p string, f os.FileInfo, err error) error {
		if path.Ext(p) != ".gpg" {
			return nil
		}

		files = append(files, p)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for i, file := range files {
		file = strings.TrimPrefix(file, passDir)
		file = strings.TrimSuffix(file, ".gpg")
		files[i] = file
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
