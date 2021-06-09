package main

// go concurrency solution is go routine, it's very light weight. can spin 10s of thousand in single program
// concurrency is defined as the number of simultaneous process
// where as parallelism means executing a single program in different threads
import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	// Sync package is not recommended for concurrency wait, Instead channels are preffered
)

func returnType(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error occured while requesting", err.Error())
		return ""
	}

	defer resp.Body.Close()

	fmt.Println(resp.Header.Get("content-type"))
	return resp.Header.Get("content-type")
}

func channels() {
	fmt.Println("Starting channels")
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		ch <- 354
	}()

	val := <-ch
	fmt.Println("Got value from channel as", val)
	fmt.Println("----")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Sending value", i, "in channel")
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Println("Got value as", val, "from channel")
	}

	fmt.Println("----")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Sending value", i, "in channel")
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println("Got the value as", i, "from the cha")
	}
}

func challenge() {
	ch := make(chan string)

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {
		go func(url string) {
			contentType := returnType(url)
			ch <- contentType
			// close(ch)
		}(url)
	}

	for content := range ch {
		fmt.Println("Got content type as", content)
	}

}

func parseSignaturesFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Problem occurred while opening a file")
	}
	defer file.Close()

	sigMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for lnum := 1; scanner.Scan(); lnum++ {
		// each line scanning
		fileds := strings.Fields(scanner.Text())
		if len(fileds) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)
		}
		sigMap[fileds[1]] = fileds[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigMap, nil
}

func fileMD5(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

type result struct {
	path  string
	match bool
	err   error
}

func md5Worker(path string, sig string, out chan *result) {
	r := &result{path: path}
	s, err := fileMD5("md5-input/" + path)

	if err != nil {
		r.err = err
		out <- r
		return
	}

	r.match = (s == sig)
	out <- r
}

func md5challenge() {
	sigMap, err := parseSignaturesFile("md5-input/md5sum.txt")
	fmt.Println("Got sigmap as", sigMap)
	if err != nil {
		fmt.Printf("error: can't read signature file - %s", err.Error())
	}

	out := make(chan *result)

	for path, hash := range sigMap {
		go md5Worker(path, hash, out)
	}

	ok := true // assuming everything passed
	for range out {
		val := <-out
		fmt.Println("Inside for loop for file", val.path)
		switch {
		case !val.match:
			fmt.Printf("Error %s: signature didn't match file %s", val.err, val.path)
			ok = false
		case val.err != nil:
			fmt.Printf("Error %s: occurred for file %s", val.err, val.path)
			ok = false
		}
		if !ok {
			break
		}
	}

	if !ok {
		fmt.Printf("Error occurred while processing the file")
		os.Exit(1)
	}

	fmt.Println("all signature passed")
}

func selectWithChannels() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	// channels can also be used along with select in case if there are any messages to be processed.
	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("----")

	out := make(chan float64)
	go func() {
		time.Sleep(200 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f from channel\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}

func GoRoutines() {
	fmt.Println("Starting go routine functions")

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {
		returnType(url)
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string) { // will span an instance of return type and move on
			// will not print anything on std out, because the program will not wait
			fmt.Println("Spawning go routine")
			returnType(url)
			wg.Done()
		}(url)
		wg.Wait()
	}

	channels()
	// challenge()
	// selectWithChannels()

	// md5challenge()

}
