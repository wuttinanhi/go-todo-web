package main

import (
	"go-todo-web/category"
	"go-todo-web/database"
	"go-todo-web/todo"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func MigrateDatabase() {
	// get database
	database := database.GetDatabase()

	// migrate database
	database.AutoMigrate(&todo.Todo{})
	database.AutoMigrate(&category.Category{})
}

func CreateMockData() {
	// get database
	db := database.GetDatabase()

	// get todo count
	var count int64
	db.Find(&todo.Todo{}).Count(&count)

	// get category count
	var categoryCount int64
	db.Find(&category.Category{}).Count(&categoryCount)

	// create category id
	var todoCategoryId, doingCategoryId, doneCategoryId uint

	// create category if not exist
	if categoryCount == 0 {
		todoCategoryId = category.CreateCategory("Todo")
		doingCategoryId = category.CreateCategory("Doing")
		doneCategoryId = category.CreateCategory("Done")
	}

	// if todo count is 0
	if count == 0 {
		// create todo
		todo.CreateTodo("Mock todo A", todoCategoryId)
		todo.CreateTodo("Mock todo B", todoCategoryId)
		todo.CreateTodo("Mock todo C", doingCategoryId)
		todo.CreateTodo("Mock todo D", doingCategoryId)
		todo.CreateTodo("Mock todo E", doneCategoryId)
	}
}

func main() {
	// migrate database
	MigrateDatabase()

	// create mock if not exist
	CreateMockData()

	// gin router
	router := gin.Default()

	// default route serve static directory
	router.Use(static.Serve("/", static.LocalFile("./frontend/public", true)))

	// todo route
	router.GET("/todo", todo.GetAllTodoRoute)           // get todo route
	router.POST("/todo/create", todo.CreateTodoRoute)   // create todo route
	router.PATCH("/todo/update", todo.UpdateTodoRoute)  // update todo route
	router.DELETE("/todo/delete", todo.DeleteTodoRoute) // delete todo route

	// category route
	router.GET("/category", category.GetAllCategoryRoute)           // get category route
	router.POST("/category/create", category.CreateCategoryRoute)   // create category route
	router.PATCH("/category/update", category.UpdateCategoryRoute)  // update category route
	router.DELETE("/category/delete", category.DeleteCategoryRoute) // delete category route

	// run router on port 3000
	router.Run(":3000")
}
