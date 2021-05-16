// Single line comment
/* Multi
line
comments
*/
package main // makes the go execute and not a package

import (
	"fmt" //format package for formating
	"strings"
)

func main() { // Function to be executed when program starts
	fmt.Println("Hello world!") // Strings are encoded in double quotes and is a unicode

	// go run file.go
	// go build welcome.go - executable file

	var x int
	var y int // assigns a zero value, for int it will be 0

	x = 1

	fmt.Printf("x=%v for object, type of variable %T\n", x, x)
	fmt.Printf("y=%v for object, type of variable %T\n", y, y)

	y = 2

	var mean int

	mean = (x + y) / 2

	fmt.Printf("result:%v, type of %T\n", mean, mean)

	mean1 := (x + 2.0) / 2.0 // this will give 1, int unless we type cast x to float64

	// = is assign and := is create and assign

	fmt.Printf("result:%v, type of %T\n", mean1, mean1)

	x, y = 2, 3 // double assignment

	fmt.Printf("x=%v for object, type of variable %T\n", x, x)
	fmt.Printf("y=%v for object, type of variable %T\n", y, y)

	// unused varibales throws an error in go

	x = 10

	if x > 5 {
		fmt.Println("X is big")
	}

	if x > 20 {
		fmt.Println("X is big")
	} else {
		fmt.Println("X is not that big")
	}

	if x > 5 && x < 15 { // logical and
		fmt.Println("X is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("X is out of range")
	}

	// if optional intialization; condition

	a := 10.0
	b := 19.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	x = 5
	switch x { // break is not required
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Default case") // print this
	}

	x = 20
	switch { // can also be used with a condition
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}

	for i := 0; i < 3; i++ { // Initialization; condition; iteration
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i <= 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	i := 0
	for i < 3 { // while loop can be achieved using this
		fmt.Println(i)
		i++
	}

	fmt.Println("----")
	i = 0
	for { // while true loop

		if i > 2 {
			break
		}

		fmt.Println(i)
		i++
	}

	// fiz buzz challenge
	fmt.Println("----")
	for i := 1; i < 21; i++ {
		switch {
		case i%(3*5) == 0:
			fmt.Println("fizz buzz")
		case i%5 == 0:
			fmt.Println("buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}

	}

	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0]) // uint 8 is a byte

	// book[0] = 116 string in go are also immutable

	fmt.Println(book[4:11])

	fmt.Println(book[4:])

	fmt.Println(book[:4])

	fmt.Println("t" + book[1:])

	fmt.Println("It was 1/2 price !")

	fmt.Println(`Printing a multiline
	poem, of the song`)

	n := 42

	val := fmt.Sprintf("%d", n)

	fmt.Printf("value is %v and type is %T\n", val, val)
	fmt.Printf("value is %q and type is %T\n", val, val)

	// for every 4 digit pair print even ended number

	/*
		count := 0
		for index1 := 1000; index1 <= 9999; index1++ {
			for index2 := index1; index2 <= 9999; index2++ {
				product := index1 * index2

				productStr := fmt.Sprintf("%d", product)
				// fmt.Printf("value is %q\n", productStr)

				if productStr[0] == productStr[len(productStr)-1] {
					count++
				}
			}
		}
		fmt.Println(count)
	*/

	loons := []string{"buggy", "clown", "tez"} // should be of same type

	fmt.Printf("loons = %v and type is %T\n  ", loons, loons)

	fmt.Println("Size of the slice is ", len(loons))

	fmt.Println("Printing the first element of the slice:", loons[1])

	fmt.Println(loons[1:])

	fmt.Println("----")
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i, name := range loons {
		fmt.Println("Index is", i, "and value is", name)
	}

	fmt.Println("----")
	for _, name := range loons {
		fmt.Println(name)
	}

	loons = append(loons, "demer")

	fmt.Println(loons)

	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)

	stocks := map[string]float64{ // all the keys should be of same type and values should be of same type
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // Must have a trailing comma or a one line decleration
	}

	fmt.Println(len(stocks))

	fmt.Println(stocks)

	fmt.Println(stocks["MSFT"])

	fmt.Println(stocks["TSLA"]) // will print 0 value of float 64

	value, ok := stocks["TSLA"] // checking if element doesn't exist

	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	delete(stocks, "AMZN")
	fmt.Println(stocks)

	fmt.Println("----")
	for key := range stocks {
		fmt.Println(key)
	}

	fmt.Println("----")
	for key, value := range stocks {
		fmt.Println("The key is", key, "and the value is", value)
	}

	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	wordMap := map[string]int{}

	for _, word := range words {
		wordMap[strings.ToLower(word)]++
	}

	fmt.Println(wordMap)

}

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"Desc"`
// 	Content string `json:"Content"`
// }

// type Articles []Article

// func returnArticles(write http.ResponseWriter,
// 	read *http.Request) {
// 	articles := Articles{
// 		Article{Title: "Test Title",
// 			Desc:    "Test Desc",
// 			Content: "Test Content"},
// 	}

// 	log.Println("Request recieved for getting all articles")
// 	json.NewEncoder(write).Encode(articles)
// }

// func returnArticlesPOST(write http.ResponseWriter, read *http.Request) {
// 	fmt.Fprintf(write, "This is POST request for articles")
// }

// func firstPage(write http.ResponseWriter,
// 	read *http.Request) {
// 	fmt.Fprintf(write, "first page loaded")
// }

// func handler() {
// 	newRouter := mux.NewRouter()

// 	newRouter.HandleFunc("/articles", returnArticles).Methods("GET")
// 	newRouter.HandleFunc("/articles", returnArticlesPOST).Methods("POST")
// 	newRouter.HandleFunc("/", firstPage)
// 	log.Fatal(http.ListenAndServe(":8081", newRouter))
// }

// func test() {
// 	fmt.Println("GO Mysql Tutorial")

// 	host := "127.0.0.1"
// 	port := 5432
// 	dbname := "raptor"
// 	password := "welcome"
// 	user := "kashif"
// 	sslmode := "disable"

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
// 	db, err := sql.Open("postgres", psqlInfo)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	insert, err := db.Query("insert into inventory (product_id, product_name) values(2, 'test2')")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer insert.Close()
// 	// fmt.Println("Closed PG connection")

// 	// ##########################

// 	fmt.Println("Registering APIs now")
// 	handler()
// }
