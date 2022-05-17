package todo

import (
	"github.com/gin-gonic/gin"
)

// todo add dto
type TodoAddDTO struct {
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
}

// todo update dto
type TodoUpdteDTO struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
}

// todo delete dto
type TodoDeleteDTO struct {
	Id int `json:"id"`
}

// serialize todo helper
func SerializeTodo(todo []Todo) []gin.H {
	// create new array
	var todos []gin.H

	// loop through todo
	for _, t := range todo {
		// add todo to array
		todos = append(todos, gin.H{
			"id":          t.ID,
			"name":        t.Name,
			"category_id": t.CategoryId,
		})
	}

	// return todo array
	return todos
}

// get all todo route
func GetAllTodoRoute(c *gin.Context) {
	// get all todo
	var todos []Todo
	db.Find(&todos)

	// send all todo json
	c.JSON(200, SerializeTodo(todos))
}

// create todo route
func CreateTodoRoute(c *gin.Context) {
	// get request body
	var dto TodoAddDTO
	c.BindJSON(&dto)

	// insert todo to database
	CreateTodo(dto.Name, dto.CategoryId)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// update todo route
func UpdateTodoRoute(c *gin.Context) {
	// get request body
	var dto TodoUpdteDTO
	c.BindJSON(&dto)

	// update todo
	UpdateTodo(dto.Id, dto.Name, dto.CategoryId)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// delete todo route
func DeleteTodoRoute(c *gin.Context) {
	// get request body
	var dto TodoDeleteDTO
	c.BindJSON(&dto)

	// delete todo
	DeleteTodo(dto.Id)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}
