package main

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	port := ":4300"
	fmt.Println("Listening on", port)
	fmt.Println("Serving static content from the /files/ directory.")
	http.HandleFunc("/files/", HandleFiles)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleFiles handles static content.
func HandleFiles(w http.ResponseWriter, r *http.Request) {
	// Get only the file name.
	Filename := path.Base(r.URL.String())
	// Whatever arbitrary logic you want. For demo purposes, we will not serve
	// content that contains the string "bad". But you could be checking that
	// the user is the owner of this file, or any additional checks you aren't
	// doing in your middleware.
	if strings.Contains(Filename, "bad") {
		fmt.Println("This file is bad and we won't serve it:", Filename)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Not authorized!"))
		return
	}
	fmt.Println("Attempting to serve", Filename)
	// Use http.ServeFile to serve the content. We don't have to worry about
	// writing not found HTTP error codes, etc., because ServeFile handles it
	// for us.
	http.ServeFile(w, r, filepath.Join(".", "files", Filename))
}
