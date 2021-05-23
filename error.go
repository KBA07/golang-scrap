package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
}

func setupLogging() {
	out, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	// Parse config files here

	return cfg, nil
}

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)

	if err != nil {
		return errors.Wrap(err, "error occured while opening pid file")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file - %s", err)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return errors.Wrap(err, "error occurred while coverting the process id, pid:")
	}

	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func ErrorFunc() {
	setupLogging()
	fmt.Println("Starting Error")

	// cfg, err := readConfig("/path/to/config.toml")
	cfg, err := readConfig("log.txt") // changed this to start the program
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		// fmt.Fprintf(os.Stderr, "%+v", err) Fprintf will write the content to the write which is passed in the first argument
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", cfg)

	// vals := []int{1, 2, 3}
	// v := vals[10]

	// fmt.Println(v)

	// file, err := os.Open("no-such-file")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	fmt.Println("File was opened")

	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)

	// use of panic is discouraged in go

	killServer("pid.txt")

}
