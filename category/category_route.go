package category

import (
	"strconv"

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
	// get category name
	name := c.PostForm("name")

	// create category
	CreateCategory(name)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// update category route
func UpdateCategoryRoute(c *gin.Context) {
	// get category id
	id, err := strconv.Atoi(c.PostForm("id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// get category name
	name := c.PostForm("name")

	// update category
	UpdateCategory(id, name)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}

// delete category route
func DeleteCategoryRoute(c *gin.Context) {
	// get category id
	id, err := strconv.Atoi(c.PostForm("id"))

	// if error return bad request
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	// delete category
	DeleteCategory(id)

	// send success json
	c.JSON(200, gin.H{"message": "success"})
}
