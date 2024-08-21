package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func validate_url(input_url string) (validated_urls []string) {
	if strings.HasPrefix(input_url, "http://") || strings.HasPrefix(input_url, "https://") {
		validated_urls = append(validated_urls, input_url)
		return validated_urls
	}
	validated_urls = append(validated_urls, "http://"+input_url)
	validated_urls = append(validated_urls, "https://"+input_url)
	return validated_urls
}

func httprun() {

	raw_urls := os.Args[1:]
	validated_urls := []string{}

	for _, url := range raw_urls {
		validated_urls = append(validated_urls, validate_url(url)...)
	}

	for _, validated_url := range validated_urls {
		fmt.Printf("validated url : %s\n", validated_url)
	}

	for _, url := range validated_urls {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			continue
			//os.Exit(1)
		}

		fmt.Printf("url response status %s: %v\n", url, resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			continue
		}
		/*
			b, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		*/
	}
}
