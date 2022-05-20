package helper

import (
	"github.com/gin-gonic/gin"
)

// JSON binding helper
func BindJSON(c *gin.Context, target interface{}) {
	if err := c.ShouldBindJSON(target); err != nil {
		panic(err)
	}
}
