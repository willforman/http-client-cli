package main

import (
	"net/http"
	"os"
	"strings"
)

func mainLoop(base Resource, projectPath string, client *http.Client) {
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
				makeRequest(req, urlSB.String(), client)
				curr = &base
				
				urlSB.Reset()
				urlSB.WriteString(base.Path)
			// select: create a resource
			case choice == len(curr.Resources) + len(curr.Requests):
				child := createResource(curr)
				curr.Resources = append(curr.Resources, child)
				writeProject(base, projectPath)
			// select: create a request
			case choice == len(curr.Resources) + len(curr.Requests) + 1:
				child := createRequest(curr)
				curr.Requests = append(curr.Requests, child)
				writeProject(base, projectPath)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("Please provide project name as argument")
	}
	projectName := os.Args[1]

	projectPath := getProjectPath(projectName)
	base := readProject(projectPath)

	client := &http.Client{}
	mainLoop(base, projectPath, client)
}
