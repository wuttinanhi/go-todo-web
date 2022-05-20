package mock

import (
	"go-todo-web/category"
	"go-todo-web/database"
	"go-todo-web/todo"
)

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
		todo.CreateTodo("Learn Go", todoCategoryId)
		todo.CreateTodo("Learn Svelte", todoCategoryId)
		todo.CreateTodo("Learn TypeScript", todoCategoryId)

		todo.CreateTodo("Write Dockerfile", doingCategoryId)
		todo.CreateTodo("Write docker-compose.yml", doingCategoryId)
		todo.CreateTodo("Create Basic API", doingCategoryId)

		todo.CreateTodo("Learn Python", doneCategoryId)
		todo.CreateTodo("Learn JavaScript", doneCategoryId)
	}
}
