package main

import (
	"fmt"
	"log"
	"os"

	// Install package - go get github.com/pelletier/go-toml
	toml "github.com/pelletier/go-toml" // explictly using toml name, _ for making package stay in import
	// Imported package will have only capital function and variable
	// go get installs latest version of sotware
	// go mod was introduced in 1.11
	// go mod init to created go.mod file for dependencies
	// will just have module decleration
	// require <package> <version> in go.mod file
	// go.mod - package you wan't, go.sum - package you actually require
	// go mod tidy - to automatically fetch and download dependies
)

type Config1 struct {
	Login struct {
		User     string
		Password string
	}
}

func ProjectManagement() {
	file, err := os.Open("config.toml")

	if err != nil {
		log.Fatalf("error: can't open config file - %s", err)
	}
	defer file.Close()

	cfg := &Config1{}
	dec := toml.NewDecoder(file)

	if err := dec.Decode(cfg); err != nil {
		log.Fatalf("error: can't decode configuration file - %s", err)
	}

	fmt.Printf("%+v\n", cfg)

}
