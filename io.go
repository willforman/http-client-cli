package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getProjectPath(projectName string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Can't read home dir")
	}

	return filepath.Join(home, ".config", "http-test", "projects", projectName + ".json")
} 

func createProject(projectPath string) Resource {
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

	base := createResource()
	writeProject(base, projectPath)
	return base
}

func readProject(projectPath string) Resource {
	fmt.Println("Reading project at " + projectPath)
	data, err := os.ReadFile(projectPath)
	if err != nil {
		return createProject(projectPath)
	}
	
	var base Resource
	err = json.Unmarshal(data, &base)
	if err != nil {
		panic(err)
	}

	return base
}

func writeProject(base Resource, projectPath string) {
	data, err := json.MarshalIndent(base, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(projectPath, data, 644)
	if err != nil {
		panic(err)
	}
}
