package main

import (
	"encoding/json"          // To encode it to json
	"fmt"                    // This package allows to format basic strings, values, or anything and print
	"github.com/gorilla/mux" // Package mux implements a request router and dispatcher.
	"log"                    // To logout errors
	"net/http"               // allows to create a server
)

type todo struct { // struct and slices has been used as databases
	ID        string `json:"id"`
	Item      string `json:"title"`
	Duration  string `json:"time"`
	Completed bool   `json:"completed"`
}

var todos = []todo{ // slice of type todo has been declared and initialised with some values
	{ID: "1", Item: "Clean Room", Duration: "10 min", Completed: false},
	{ID: "2", Item: "Brush Teeth", Duration: "10 min", Completed: false},
	{ID: "3", Item: "Go For a Walk", Duration: "15 min", Completed: false},
	{ID: "4", Item: "Breakfast", Duration: "20 min", Completed: false},
	{ID: "5", Item: "Study", Duration: "1 hour", Completed: false},
	{ID: "6", Item: "Attend Meeting", Duration: "30 min", Completed: false},
}

// @GET ('/todos') ...to fetch data
func getTodos(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// @DELETE ('/todos/{id}') ...to delete a value by using id
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//The name mux stands for "HTTP request multiplexer". 
	//Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	params := mux.Vars(r)
	//loop over the todos, range
	//delete the todo with the id that you've sent
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

// @GET method ('/todos/{id}') ...to fetch data by id
func getTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// @POST ('/todos') ...to create data
func createTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	var todo todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// @PUT ('/todos/{id}') ...update data
func updateTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the todos, range
	//delete the todo with the id that you've sent
	//add a new todo - the todo that we send in the body of our json content inside Postman/ Thunder Client
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var todo todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
		}
	}
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodos).Methods("GET")           // func to display all the data present in the todo[]
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")       // func to get a single value
	r.HandleFunc("/todos", createTodo).Methods("POST")        // func to create and add a new value in the existing data
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")    // func to update an existing value
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE") // func to delete a value from the data

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r)) // to start a server at port:8080
}
