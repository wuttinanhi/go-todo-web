package category

import (
	"go-todo-web/database"

	"gorm.io/gorm"
)

var db *gorm.DB = database.GetDatabase()

// category struct
type Category struct {
	gorm.Model
	Name string `json:"name"`
}

/*
	Create catregory
	@param name string
*/
func CreateCategory(name string) uint {
	// create new category
	category := Category{Name: name}

	// save category
	db.Create(&category)

	// return id
	return category.ID
}

/*
	Update category
	@param id int
	@param name string
*/
func UpdateCategory(id int, name string) {
	// update category
	db.Model(&Category{}).Where("id = ?", id).Update("name", name)
}

/*
	Delete category
	@param id int
*/
func DeleteCategory(id int) {
	// delete category
	db.Delete(&Category{}, id)
}

/*
	Check valid category ID
	@param id int
*/
func CheckValidCategoryId(id uint) bool {
	// get category
	var category Category
	db.First(&category, id)

	// check if category is valid
	return category.ID != 0
}
