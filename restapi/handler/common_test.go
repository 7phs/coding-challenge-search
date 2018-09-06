package handler

import "github.com/gin-gonic/gin"

func testGinMode() func() {
	mode := gin.Mode()
	gin.SetMode(gin.TestMode)

	return func() {
		gin.SetMode(mode)
	}
}
