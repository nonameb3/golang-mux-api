package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nonameb3/golang-mux-api/controllers"
	"github.com/nonameb3/golang-mux-api/model"
)

func main() {
	// inital route
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	model.Books = append(model.Books, model.Book{ID: "1", Title: "Book One", Author: &model.Author{Firstname: "John", Lastname: "Doe"}})
	model.Books = append(model.Books, model.Book{ID: "2", Title: "Book Two", Author: &model.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/", controllers.SetString).Methods("GET")
	r.HandleFunc("/api/json", controllers.GetJSON).Methods("GET")
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{ID}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{ID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{ID}", controllers.DeleteBook).Methods("DELETE")

	// Start server
	port := ":10000"
	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// json
// {
// 	"id": "999",
// 	"title": "new books",
// 	"author" : {
// 			"firstname": "waraphon",
// 			"lastname": "roonnapai"
// 	}
// }
