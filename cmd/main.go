package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	searcherPkg "pulley.com/shakesearch/pkg/searcher"
	shakesearchPkg "pulley.com/shakesearch/pkg/searcher/shakesearch"
	filestorePkg "pulley.com/shakesearch/pkg/datastore/filestore"
	datastorerPkg "pulley.com/shakesearch/pkg/datastore"
	// searcherIntr "pulley.com/shakesearch/pkg/searcher"
)

func main() {
	searcher := shakesearchPkg.Searcher{}
	// err := searcher.Load("completeworks.txt")
	fileStore, err := filestorePkg.NewFileStore("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// this is where we are currently planning to pass search algo and storage
	// we can swap the search algo and storage as required later.
	http.HandleFunc("/search", handleSearch(searcher, fileStore))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSearch(searcher searcherPkg.Searcher, store datastorerPkg.Storer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.Search(query[0], store)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}
