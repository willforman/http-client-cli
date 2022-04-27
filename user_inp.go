package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Request struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Method string `json:"method"`
	Data map[string]interface{}  `json:"data"`
}

type Resource struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Resources []Resource `json:"resources"`
	Requests []Request `json:"requests"`
}

func getChoice(numChoices int) int {
	for {
		choice := numChoices
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("choice is not an integer")
			continue
		}
		
		if choice >= numChoices || choice < 0 {
			fmt.Println("choice is out of range")
			continue
		}

		return choice
	}
}

func createResource(reader *bufio.Reader) Resource {
	println("Name: ");
	name, err := reader.ReadString('\n');
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSuffix(name, "\n")

	println("Path: ");
	path, err := reader.ReadString('\n');
	if err != nil {
		fmt.Println(err)
	}
	path = strings.TrimSuffix(path, "\n")

	return Resource{
		Name: name,
		Path: path,
		Resources: []Resource{},
		Requests: []Request{},
	}
}

func createRequest(reader *bufio.Reader) Request {
	println("Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSuffix(name, "\n")

	println("Path: (default = '/')")
	path, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	path = strings.TrimSuffix(path, "\n")

	var method string
	println("Method:\n0: GET\n1: POST\n2: PUT\n3: DELETE")
	methodChoice := getChoice(4)
	switch methodChoice {
		case 0:
			method = "GET"
		case 1:
			method = "POST"
		case 2:
			method = "PUT"
		case 3:
			method = "DELETE"
	}

	var data map[string]interface{}
	println("Data: ")
	err = json.NewDecoder(os.Stdin).Decode(&data)
	if err != nil {
		panic(err)
	}

	return Request{
		Name: name,
		Path: path,
		Method: method,
		Data: data,
	}
}

func getDataStr(req Request) string {
	data, err := json.MarshalIndent(req.Data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func createProject(reader *bufio.Reader, projectPath string) Resource {
	fmt.Print("No project found. Create one [y/N]? ")
	var choice rune
	_, err := fmt.Scanf("%c", &choice)
	if err != nil {
		fmt.Println(err)
		panic(err);
	}

	if (choice != 'y') {
		return Resource{}
	}

	base := createResource(reader)
	writeProject(base, projectPath)
	return base
}
