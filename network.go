package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func makeRequest(req Request, urlStr string, baseUrl *url.URL, client *http.Client) {
	println(req.Name)
	println(req.Method + " " + urlStr)

	var httpReq *http.Request
	var err error

	if (req.Method == "GET") {
		httpReq, err = http.NewRequest(req.Method, urlStr, nil)
	} else {
		dataStr := getDataStr(req)
		reader := strings.NewReader(dataStr)

		httpReq, err = http.NewRequest(
			req.Method,
			urlStr, 
			reader,
		)
		httpReq.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.Header.Get("Content-Type") == "application/json; charset=utf-8" {
		var bodyIndented bytes.Buffer
		err = json.Indent(&bodyIndented, body, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(bodyIndented.String())
	} else {
		fmt.Println(string(body))
	}

	// We want to save any cookies we get from the server
	if len(resp.Cookies()) != 0 {
		fmt.Println("Saving following cookies: ")
		for _, cookie := range resp.Cookies() {
			fmt.Println(cookie)
		}
		client.Jar.SetCookies(baseUrl, resp.Cookies()) 
	}
}
