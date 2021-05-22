package main

import (
	"fmt"
	"math"
	"net/http"
)

func add(a int, b int) int { // can also have merged declearation a , b int
	return a + b
}

func divmod(a int, b int) (int, int) { // return types
	return a / b, a % b // go functions can have more than one return
}

func doubleAt(values []int, i int) { // slices are stored in heap, so that means modifying the instance will change the original slice
	values[i] *= 2
}

func double(n int) {
	n *= 2 // since n lifetime is bound to that of this function, hence the value is lost and not reflected to the original variable
}

func doublePtr(n *int) {
	*n *= 2
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("square root of negative value tried %f", n)
	}
	return math.Sqrt(n), nil
}

func cleanup(name string) {
	fmt.Println("Cleaning up:", name)
}

func worker() {
	defer cleanup("A") // defer will cause the function to execute at last, similar to finally of python
	defer cleanup("B") // will be called first, works on the principle of last in first out

	fmt.Println("Worker")
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil
}

func Functions() {
	addVal := add(3, 6)

	div, mod := divmod(6, 3)

	fmt.Println("Result of addition is", addVal)
	fmt.Printf("Result if division is: %d, result of mod is: %d", div, mod)

	value := []int{1, 2, 3, 4}
	doubleAt(value, 2) // passed by address, same for slice as well
	fmt.Println(value)

	val := 2
	double(val) // passed by value
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)

	s1, err := sqrt(-1)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(3)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
	} else {
		fmt.Println(s2)
	}

	worker()

	contenttype, err := contentType("https://google.com")

	if err != nil {
		fmt.Println("Error occurred while requesting", err.Error())
	} else {
		fmt.Println(contenttype)
	}

}
