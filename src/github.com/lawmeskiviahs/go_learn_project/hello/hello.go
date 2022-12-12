// package main

// import (
//     "fmt"
//     "log"
//     "example.com/greetings"
//     "example.com/maths"
//     "example/localTest" //modules cannot have underscores in their names or paths
//     "github.com/gorilla/mux"
// )

// func main() {

//     log.SetPrefix("greetings: ") // sets prefix to logs
//     log.SetFlags(0) // without this date time and line number are also printed in logs

//     // Get a greeting message and print it.
//     message, err := greetings.Hello("hello")
//     if err != nil {
//         log.Fatal(err) // exits the code with the err
//     }
//     fmt.Println(message)
// 	sum := maths.Add(2, 3)
// 	fmt.Println(sum)
//     value:=localTest.Export()
//     fmt.Println(value)
//     r:=mux.NewRoutera()
//     fmt.Println(r)

// }

// CRUD
// package main

// import ("fmt"
// "log"
// "encoding/json"
// // "math/rand"
// 	"net/http"
// 	// "strconv"
// 	"github.com/gorilla/mux"
// )

// type Movie struct{
// 	Id string `json:"id"`
// 	Title string `json:"title"`
// }

// var movies []Movie

// func getMovies(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(movies)
// }

// func main(){
// 	movies=append(movies, Movie{"1","IronMan"})
// 	movies=append(movies, Movie{"2","Thor"})
// 	r:=mux.NewRouter()
// 	r.HandleFunc("/movies",getMovies).Methods("GET")
// 	// r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
// 	// r.HandleFunc("/movies",createMovie).Methods("POST")
// 	// r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
// 	// r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

// 	fmt.Println("Starting server at port 8000")
// 	log.Fatal(http.ListenAndServe(":8000",r))
// }


//channels
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazo n.com",
	}
	// links = append(links, "http://google.com")
	// links = append(links, "http://facebook.com")
	// links = append(links, "http://stackoverflow.com")
	// links = append(links, "http://golang.org")
	// links = append(links, "http://amazon.com")

	for _, link := range links {
		checkLink(link)
	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link ," site down")
		return
	}
	fmt.Println(link, " site fine")
}
