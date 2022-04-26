package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func makeRequest(req Request, url string, client *http.Client) {
	println(req.Name)
	println(req.Method + " " + url)

	var httpReq *http.Request
	var err error

	if (req.Method == "GET") {
		httpReq, err = http.NewRequest(req.Method, url, nil)
	} else {
			dataStr := getDataStr(req)
			reader := strings.NewReader(dataStr)

			httpReq, err = http.NewRequest(
				req.Method,
				url, 
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
	fmt.Println(string(body))
}
