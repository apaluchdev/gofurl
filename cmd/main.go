package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"example.com/link"
)

func main() {
	sourceURI := ""

	if len(os.Args) < 2 {
		sourceURI = "../htmlTests/ex4.html"
	} else {
		sourceURI = os.Args[1]
	}

	source := io.ReadCloser(nil)

	// Check if the input is a URL or a local file
	if strings.HasPrefix(sourceURI, "http") {
		response, err := http.Get(sourceURI)
		if err != nil {
			log.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Failed to fetch URL: %s", response.Status)
		}
		source = response.Body

	} else {
		file, err := os.Open(sourceURI)
		if err != nil {
			log.Fatal(err)
		}
		source = file
	}

	var links, _ = link.ParseURLs(source)

	for _, l := range links {
		fmt.Printf("Href: %v, Link Text: %v\n", l.Href, l.Text)
	}

	fmt.Println("Done")
}
