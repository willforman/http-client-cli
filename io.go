package main

import (
	"encoding/json"
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

func readProject(projectPath string) Resource {
	data, err := os.ReadFile(projectPath)
	if err != nil {
		panic("Project doesn't exist")
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
