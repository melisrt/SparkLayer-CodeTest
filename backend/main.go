package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ToDo struct {       //Defines a struct named ToDo
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`    //"urgent", "medium", "low"
	Completed   bool   `json:"completed"`
}

var (
	todos []ToDo        //In-memory storage for todos
	nextID int = 1      //Counter for generating unique IDs

	//Possible future TODO for practice and learning: Adding a sync.Mutex to protect todos and nextID from concurrent access.

)

//validateAndSetPriority validates and sets default priority value.
//If priority is empty defaults to "medium".
func validateAndSetPriority(priority string) (string, error) {
	if priority == "" {
		priority = "medium"
	}
	if priority != "urgent" && priority != "medium" && priority != "low" {
		return "", fmt.Errorf("priority must be 'urgent', 'medium', or 'low'")
	}
	return priority, nil
}

func main() {
	http.HandleFunc("/", ToDoListHandler)
	http.ListenAndServe(":8080", nil)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//Handle GET request: return all todos
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
			return
		}
		return
	}


	//Handle POST request: create a new todo
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()
		
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
		
		//Validate and set priority
		priority, err := validateAndSetPriority(newTodo.Priority)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newTodo.Priority = priority
		
		//Set default values for new todos
		newTodo.ID = nextID
		nextID++
		newTodo.Completed = false

		todos = append(todos, newTodo)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(newTodo); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	//Handle DELETE request: delete a todo by ID
	if r.Method == http.MethodDelete {
		//Get ID from query parameter
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}
		
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		
		//find and remove the todo
		found := false
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				found = true
				break
			}
		}
		
		if !found {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Todo deleted successfully"))
		return
	}

	//Handle PUT request: update a todo by ID
	if r.Method == http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()
		
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}
		
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		
		//decode the updated todo from request body
		var updatedTodo ToDo
		err = json.NewDecoder(r.Body).Decode(&updatedTodo)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		
		//Validate required fields
		if updatedTodo.Title == "" || updatedTodo.Description == "" {
			http.Error(w, "Title and description are required", http.StatusBadRequest)
			return
		}
		
		//Vvalidate and set priority
		priority, err := validateAndSetPriority(updatedTodo.Priority)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedTodo.Priority = priority
		
		//Find and update the todo
		found := false
		for i, todo := range todos {
			if todo.ID == id {
				//Update the todo 
				updatedTodo.ID = id
				todos[i] = updatedTodo
				found = true
				break
			}
		}
		
		if !found {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(updatedTodo); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}
