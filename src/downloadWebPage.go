package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func runDownloadWebPageDemo() {
	url := "https://www.google.com"
	fmt.Printf("Downloading page %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to download page %v - ", url), err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to download page %v - ", url), err)
	}
}
