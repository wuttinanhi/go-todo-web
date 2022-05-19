package todo

import (
	"go-todo-web/category"
	"go-todo-web/database"

	"gorm.io/gorm"
)

var db *gorm.DB = database.GetDatabase()

// todo struct
type Todo struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
}

/*
	Create todo
	@param name string
*/
func CreateTodo(name string, categoryId uint) uint {
	// check if category is valid
	if category.CheckValidCategoryId(categoryId) == false {
		panic("category id is not valid")
	}

	// create new todo
	todo := Todo{Name: name, CategoryId: categoryId}

	// save todo
	db.Create(&todo)

	// return id
	return todo.ID
}

/*
	Update todo
	@param id int
	@param name string
*/
func UpdateTodo(id int, name string, categoryId uint) {
	// check if category is valid
	if category.CheckValidCategoryId(categoryId) == false {
		panic("category id is not valid")
	}

	// update todo
	db.Model(&Todo{}).Where("id = ?", id).Update("name", name).Update("category_id", categoryId)
}

/*
	Delete todo
	@param id int
*/
func DeleteTodo(id int) {
	// delete todo
	db.Delete(&Todo{}, id)
}
