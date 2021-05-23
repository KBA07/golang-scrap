// Single line comment
/* Multi
line
comments
*/
package main // makes the go execute and not a package

func main() {
	Basic()
	Functions()
	Oop()
	ErrorFunc()
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
