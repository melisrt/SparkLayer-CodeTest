package main

import (
	"encoding/json"
	"net/http"
)

type ToDo struct {       //Defines a struct named ToDo
	Title string `json:"title"`
	Description string `json:"description"`
}

var todos []ToDo        //In-memory storage for todos

func main() {
	http.HandleFunc("/", ToDoListHandler)
	http.ListenAndServe(":8080", nil)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//Handle GET request: return all todos
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
		return
	}


	//Handle POST request: create a new todo
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		var newTodo ToDo
		err := json.NewDecoder(r.Body).Decode(&newTodo)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		if newTodo.Title == "" || newTodo.Description == "" {
			http.Error(w, "Title and description are required", http.StatusBadRequest)
			return
		}

		todos = append(todos, newTodo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newTodo)
		return

	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}
