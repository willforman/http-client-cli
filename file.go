package main

import (
	"bufio"
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


func readProject(reader *bufio.Reader, projectPath string) Resource {
	fmt.Println("Reading project at " + projectPath)
	data, err := os.ReadFile(projectPath)
	if err != nil {
		return createProject(reader, projectPath)
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
	err = os.WriteFile(projectPath, data, 0644)
	if err != nil {
		panic(err)
	}
}
