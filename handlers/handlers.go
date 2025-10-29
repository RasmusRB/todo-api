package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	Detail string `json:"detail"`
}

// @Summary Get all todos
// @Description Get a list of all todo items
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} ToDo
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []ToDo{
		{ID: "1", Title: "Buy groceries", Done: false, Detail: "Milk, Bread, Eggs"},
		{ID: "2", Title: "Read book", Done: true, Detail: "The Go Programming Language"},
		{ID: "3", Title: "Exercise", Done: false, Detail: "30 minutes of running"},
	})
}

// @Summary Get a single todo
// @Description Get a todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} ToDo
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	// Sample data - in real app, this would query a database
	todos := []ToDo{
		{ID: "1", Title: "Buy groceries", Done: false, Detail: "Milk, Bread, Eggs"},
		{ID: "2", Title: "Read book", Done: true, Detail: "The Go Programming Language"},
		{ID: "3", Title: "Exercise", Done: false, Detail: "30 minutes of running"},
	}

	// Find todo by ID
	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	// Not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// @Summary Create a new todo
// @Description Create a new todo item from JSON body
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body ToDo true "Todo object"
// @Success 201 {string} string "Todo ID"
// @Failure 400 {object} map[string]string
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var newTodo ToDo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo.ID)
}

// @Summary Update a todo
// @Description Update an existing todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body ToDo true "Updated Todo object"
// @Success 200 {object} ToDo
// @Router /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo ToDo
	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}
	updatedTodo.ID = id
	c.IndentedJSON(http.StatusOK, updatedTodo)
}

// @Summary Delete a todo
// @Description Delete a todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204 {string} string "No Content"
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
