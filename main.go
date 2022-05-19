package main

import (
	"embed"
	"go-todo-web/category"
	"go-todo-web/database"
	"go-todo-web/todo"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//go:embed frontend/public
var embedFS embed.FS

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// check if indexing is allowed
	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string, index bool) static.ServeFileSystem {
	subFS, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(subFS),
		indexes:    index,
	}
}

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

	//apply CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// handle panic
	router.Use(gin.CustomRecoveryWithWriter(nil, func(c *gin.Context, err interface{}) {
		// if err is ValidationErrors
		if err, ok := err.(validator.ValidationErrors); ok {
			// error string
			var errorString string = err.Error()

			// split error string at \n
			var errors []string = strings.Split(errorString, "\n")

			// send error json
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}

		// return internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal Server Error"})
	}))

	// default route serve static directory
	router.Use(static.Serve("/", EmbedFolder(embedFS, "frontend/public", true)))

	// todo route
	router.GET("/api/todo", todo.GetAllTodoRoute)           // get todo route
	router.POST("/api/todo/create", todo.CreateTodoRoute)   // create todo route
	router.PATCH("/api/todo/update", todo.UpdateTodoRoute)  // update todo route
	router.DELETE("/api/todo/delete", todo.DeleteTodoRoute) // delete todo route

	// category route
	router.GET("/api/category", category.GetAllCategoryRoute)           // get category route
	router.POST("/api/category/create", category.CreateCategoryRoute)   // create category route
	router.PATCH("/api/category/update", category.UpdateCategoryRoute)  // update category route
	router.DELETE("/api/category/delete", category.DeleteCategoryRoute) // delete category route

	// run router on port 3000
	router.Run(":3000")
}
