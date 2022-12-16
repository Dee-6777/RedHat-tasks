//The name mux stands for "HTTP request multiplexer".
//Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.

package main

import (
	"encoding/json" // To encode it to json
	"fmt"           // This package allows to format basic strings, values, or anything and print
	"log"           // To logout errors
	"net/http"      // allows to create a server

	"github.com/gorilla/mux" // gorilla/mux is a package which adapts to Go’s default HTTP router and establish different routes
)

/*...............................model.....................................*/

type todo struct { // struct and slices has been used as database
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

/*................................ controllers ..................................................*/

// @GET ('/todos') ...to fetch data
func getTodos(w http.ResponseWriter, r *http.Request) { // r is the pointer of the request that you'll send from thunder client/postman and w is the response from the function
	//set json content type
	w.Header().Set("Content-Type", "application/json") //our golang struct needs to be converted to json coming into it so it sends the http json header to the browser to inform it what kind of data it expects.
	json.NewEncoder(w).Encode(todos)                   // the response we're sending is encoded to json (whole todos is converted to json)
}

/* In an http.HandlerFunc, the http.ResponseWriter value (named w in your handlers) is used to control the response information being written back to the client that made the request, such as the body of the response or the status code.
Then, the *http.Request value (named r in your handlers) is used to get information about the request that came into the server, such as the body being sent in the case of a POST request*/

// @DELETE ('/todos/{id}') ...to delete a value by using id
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	//set content type indicates that the request body format is in JSON
	w.Header().Set("Content-Type", "application/json")
	// we are declaring a variable called params which will store the request
	params := mux.Vars(r)
	//loop over the todos
	//delete the todo with the id that you've sent
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...) // inplace of todos[:index] it will append all other todos after that
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
			json.NewEncoder(w).Encode(item) //will convert to json the only value who's id has been matched
			return
		}
	}
}

// @POST ('/todos') ...to create data
func createTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	var todo todo
	_ = json.NewDecoder(r.Body).Decode(&todo) //Try to decode the entire request body into struct we don't want to store the error so used a blank identifier which will respond to the client with the error message and a 400 status code.
	todos = append(todos, todo)               // appended the new todo inside todos
	json.NewEncoder(w).Encode(todo)           // again encoded to json
}

// @PUT ('/todos/{id}') ...update data
func updateTodo(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params will store the request
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
	//First create a new request router.
	//The router is the main router for your web application and will later be passed as parameter to the server.
	//It will receive all HTTP connections and pass it on to the request handlers you will register on it.
	r := mux.NewRouter()

	/*..............................Routers............................*/

	// Registering Request Handlers
	// if an upcoming request URL matches one of the paths, the correspondng handler is called passing (http.ResponseWriter, *http.Request) as parameters
	r.HandleFunc("/todos", getTodos).Methods("GET")           // Registering get request handler for fetching all the elements present in our database
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")       // Registering get request handler which will fetch a single element using it's id in our database
	r.HandleFunc("/todos", createTodo).Methods("POST")        // Registering post request handler to create and add a new element in our database
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")    // Registering put request handler to update a todo
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE") // Registering delete request handler to delete a todo

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r)) // to start a server at port:8080
}
