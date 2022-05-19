package category

import (
	"go-todo-web/helper"

	"github.com/gin-gonic/gin"
)

// get category route
func GetAllCategoryRoute(c *gin.Context) {
	// get all category
	var categories []Category
	db.Find(&categories)

	// return category json
	c.JSON(200, categories)
}

// create category route
func CreateCategoryRoute(c *gin.Context) {
	// parse create category dto
	var dto CategoryAddDTO
	helper.BindJSON(c, &dto)

	// create category
	CreateCategory(dto.Name)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// update category route
func UpdateCategoryRoute(c *gin.Context) {
	// parse update category dto
	var dto CategoryUpdateDTO
	helper.BindJSON(c, &dto)

	// update category
	UpdateCategory(dto.Id, dto.Name)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// delete category route
func DeleteCategoryRoute(c *gin.Context) {
	// parse delete category dto
	var dto CategoryDeleteDTO
	helper.BindJSON(c, &dto)

	// delete category
	DeleteCategory(dto.Id)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}
