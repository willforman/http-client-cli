package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func createResource(parent *Resource) Resource {
	var name string
	println("Name: ");
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}

	var path string
	println("Path: ");
	_, err = fmt.Scan(&path)
	if err != nil {
		fmt.Println(err)
	}

	return Resource{
		Name: name,
		Path: path,
		Resources: []Resource{},
		Requests: []Request{},
	}
}

func createRequest(parent *Resource) Request {
	var name string
	println("Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}

	var path string
	println("Path: (default = '/')")
	_, err = fmt.Scan(&path)
	if err != nil {
		fmt.Println(err)
	}

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


