package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// Route Handler getTodos
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// Route Handler addTodo
func addTodo(context *gin.Context) {
	var newTodo todo
	
	// Potential error if json does not fit the struct types
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

// Route Handler getTodo
func getTodo(context *gin.Context) {
	id := context.Param("id") // Comes from /todos/:id
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo);
}

// Route Handler toggleTodo
func toggleTodo(context *gin.Context) {
	id := context.Param("id") // Comes from /todos/:id
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo);
}

func getTodoById(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func main() {	
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodo)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}