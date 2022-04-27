package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please provide project name as argument")
	}
	projectName := os.Args[1]

	projectPath := getProjectPath(projectName)
	reader := bufio.NewReader(os.Stdin)
	base := readProject(reader, projectPath)

	if base.Name == "" {
		return
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
	}
	baseUrl, err := url.Parse(base.Path)
	if err != nil {
		panic(err)
	}

	mainLoop(base, projectPath, client, reader, baseUrl)
}

func mainLoop(base Resource, projectPath string, client *http.Client, reader *bufio.Reader, baseUrl *url.URL) {
	curr := &base
	var urlSB strings.Builder
	urlSB.WriteString(base.Path)
	for {
		printResource(curr)
		choice := getChoice(len(curr.Resources) + len(curr.Requests) + 2)
		switch {
			// select resource
			case choice < len(curr.Resources):
				curr = &curr.Resources[choice]
				urlSB.WriteString(curr.Path)
			// select: request
			case choice < len(curr.Resources) + len(curr.Requests):
				req := curr.Requests[choice]
				urlSB.WriteString(req.Path)
				makeRequest(req, urlSB.String(), baseUrl, client)
				curr = &base
				
				urlSB.Reset()
				urlSB.WriteString(base.Path)
				fmt.Println("==================================================")
			// select: create a resource
			case choice == len(curr.Resources) + len(curr.Requests):
				child := createResource(reader)
				curr.Resources = append(curr.Resources, child)
				writeProject(base, projectPath)
			// select: create a request
			case choice == len(curr.Resources) + len(curr.Requests) + 1:
				child := createRequest(reader)
				curr.Requests = append(curr.Requests, child)
				writeProject(base, projectPath)
		}
	}
}

func printResource(base *Resource) {
	fmt.Println(base.Name)
	
	for idx, subResource := range base.Resources {
		fmt.Printf("%d: %s\n", idx, subResource.Name)
	} 
	for idx, requests := range base.Requests {
		fmt.Printf("%d: %s\n", idx + len(base.Resources), requests.Name)
	} 
	fmt.Printf("%d: Add sub-resource\n", len(base.Resources) + len(base.Requests))
	fmt.Printf("%d: Add sub-request\n", len(base.Resources) + len(base.Requests) + 1)
}
