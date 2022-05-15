package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	// get todo name
	name := c.PostForm("name")
	// get todo category id
	categoryId, err := strconv.Atoi(c.PostForm("category_id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// create todo
	CreateTodo(name, uint(categoryId))

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// update todo route
func UpdateTodoRoute(c *gin.Context) {
	// get todo id
	id, err := strconv.Atoi(c.PostForm("id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// get todo name
	name := c.PostForm("name")
	// get todo category id
	categoryId, err := strconv.Atoi(c.PostForm("category_id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// update todo
	UpdateTodo(id, name, uint(categoryId))

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// delete todo route
func DeleteTodoRoute(c *gin.Context) {
	// get todo id as int
	id, err := strconv.Atoi(c.PostForm("id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// delete todo
	DeleteTodo(id)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}
