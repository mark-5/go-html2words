package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/http"
	"os"
)

var (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
)

func printWords (body io.ReadCloser) error {
	last := atom.Html
	z    := html.NewTokenizer(body)

	for {
		tt := z.Next()
		token := z.Token()
		tag   := token.DataAtom
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			if ( last != atom.Script   ) &&
			   ( last != atom.Style    ) &&
			   ( last != atom.Noscript ) {
				fmt.Printf("%s", token.Data)
			}
		case html.StartTagToken:
			if tag == atom.P {
				fmt.Print("\t")
			}
			last = tag
		case html.EndTagToken:
			if tag == atom.P {
				fmt.Print("\n\n")
			}
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s URL\n", os.Args[0])
		os.Exit(2)
	}
	client := &http.Client{}
	url    := os.Args[1]

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	err = printWords(resp.Body)
	switch err {
	case nil, io.EOF:
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
