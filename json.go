// http test package for automating http test
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var data = `
{
	"user": "kashif",
	"type": "deposit",
	"amount": 10000.3
}
`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func JSON() {
	rdr := bytes.NewBufferString(data)
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error occurred while read -%s", err)
	}

	fmt.Printf("got: %+v\n", req)

	prevBalance := 8500.0

	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error can't encode the json %s", err)
	}

	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)

	job := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer
	encPost := json.NewEncoder(&buf)

	if err := encPost.Encode(job); err != nil {
		log.Fatalf("error occurred while encoding %s", err)
	}

	responsePost, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't request a post API")
	}
	defer responsePost.Body.Close()

	io.Copy(os.Stdout, responsePost.Body)
}
