package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Trade struct { // user defined data type
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if len(symbol) == 0 {
		return nil, fmt.Errorf("symbol can't be empty")
	}

	if volume == 0 {
		return nil, fmt.Errorf("volume can't be empty")
	}

	if price == 0 {
		return nil, fmt.Errorf("price can't be empty")
	}

	return &Trade{symbol, volume, price, buy}, nil
}

// Constructor

// define functions to that struct
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price

	if t.Buy {
		value = -value
	}

	return value
}

type Point struct {
	X int
	Y int
}

func (p Point) Move(x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) MovePtr(x, y int) {
	p.X += x
	p.Y += y
}

type Square struct {
	Centre Point
	Length int
}

func NewSquare(x, y, length int) (*Square, error) {
	if length == 0 {
		return nil, fmt.Errorf("length of square can't be zero")
	}
	return &Square{Point{x, y}, length}, nil
}

func (s *Square) Move(x, y int) {
	s.Centre.MovePtr(x, y)
}

func (s *Square) Area() float64 {
	return float64(s.Length * s.Length)
}

// write common slice which takes two objects and prints its Area

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface { // This is widely used in io.Writer functions
	Area() float64 // this is required and should be same for both the objects
}

// Write a struct called Capper that has a filed to another io.Writer and transforms everythin to uppercase.
// Capper should implement io.Writer

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(data []byte) (n int, err error) { // standard interface
	offset := byte('a' - 'A')
	fmt.Println("offset is:", offset)

	for i, char := range data {
		if char >= 'a' && char <= 'z' {
			data[i] -= offset
		}
	}
	// data = []byte(strings.ToUpper(string(data)))
	return c.wtr.Write(data)
}

func Oop() {
	t1 := Trade{"MSFT", 10, 99.98, true} // order of declearation to be followed

	fmt.Println(t1)
	fmt.Printf("%+v\n", t1) // this will print the object with named attributes

	fmt.Println(t1.Symbol)

	t2 := Trade{ // Order of symbol can't be followed
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t2)

	t3 := Trade{} // will initialize all the values with zero fields
	fmt.Printf("%+v\n", t3)

	// Everything starting with Capital is public rest is private

	fmt.Println(t2.Value()) // accessing the value with the functions

	p := Point{1, 2} // initializing the struct
	fmt.Printf("%+v\n", p)

	p.Move(3, 4) // this move was performed locally and didn't affect the objects identity
	fmt.Printf("%+v\n", p)

	p.MovePtr(3, 4)
	fmt.Printf("%+v\n", p)

	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("Error occurred can't create the trade: %s\n", err)
	}

	fmt.Printf("%+v\n", t)

	sq, err := NewSquare(0, 0, 1)

	if err != nil {
		// log.Fatalf("Will print and error and exit the program with non success status code")
		fmt.Printf("Error occurred %s \n", err)
	}

	fmt.Printf("%+v\n", sq)

	fmt.Println(sq.Area())

	sq.Move(2, 2)
	fmt.Printf("%+v\n", sq)

	// making a slice of shape

	shapes := []Shape{}

	shapes = append(shapes, sq)

	circle := Circle{2}

	shapes = append(shapes, circle)

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}

	// open a file interface

	// file, err := os.OpenFile("file.txt", os.O_RDWR, 0777)
	c := &Capper{os.Stdout}

	fmt.Fprintln(c, "Hello there")

}
