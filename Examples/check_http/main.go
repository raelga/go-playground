package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	// flag parsing
	url := flag.String("url", "", "URL to check.")
	content := flag.String("content", "", "Content to match.")
	code := flag.Int("code", 0, "HTTP status response to match.")
	flag.Parse()

	// parameter validation
	if *url == "" {
		// print stderr
		fmt.Fprintln(os.Stderr, "Missing url.")
		// return command usage using flag
		flag.Usage()
		// return custom exit code
		os.Exit(1)
	}

	// call a function, retrieve result and error
	result, err := checkURL(*url, *content, *code)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// switch statement
	switch result {
	case true:
		// if / else if / else blocks
		if *code == 0 {
			fmt.Println("OK - URL ", *url, "contains", *content)
		} else if *content == "" {
			fmt.Println("OK - URL", *url, "responds with status", *code)
		} else {
			fmt.Println("OK - URL ", *url, "responds with status", *code, " and contains ", *content)
		}
	case false:
		fmt.Println("CRITICAL - URL didn't match any condition.")
	}

}

// return vars defined at func level
func checkURL(url string, content string, code int) (result bool, err error) {

	if content != "" {

		result, err = checkURLContent(url, content)

		if err != nil {
			return false, err
		}

	}

	if code != 0 {

		result, err = checlURLCode(url, code)

		if err != nil {
			return false, err
		}

	}

	return result, nil
}

func checkURLContent(url string, content string) (bool, error) {

	// htt client
	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	// defer execution
	defer response.Body.Close()

	// read url body into byte[] using ioutil
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// use type functions
	return strings.Contains(string(body), content), nil

}

func checlURLCode(url string, code int) (bool, error) {

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}

	return response.StatusCode == code, nil

}
