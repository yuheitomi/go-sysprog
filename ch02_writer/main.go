package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "world",
	}

	gzipWrite(w, source)
}

func gzipWrite(w io.Writer, source map[string]string) {
	gw := gzip.NewWriter(w)
	mw := io.MultiWriter(gw, os.Stdout)

	gw.Header.Name = "test.json"
	if err := json.NewEncoder(mw).Encode(source); err != nil {
		log.Fatal(err)
	}
	gw.Flush()
}
