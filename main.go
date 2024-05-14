package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Only GET method is allowed for the /new endpoint.
	router.Any(NEW_GAME, func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			newHandler(context)
		default:
			context.JSON(http.StatusMethodNotAllowed, gin.H{ERROR: METHOD_NOT_ALLOWED}) // 405
		}
	})

	// Only POST method is allowed for the /validate endpoint.
	router.Any(VALIDATE, func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodPost:
			validateHandler(context)
		default:
			context.JSON(http.StatusMethodNotAllowed, gin.H{ERROR: METHOD_NOT_ALLOWED}) // 405
		}
	})

	fmt.Println("Server is running on", SERVER)
	router.Run(SERVER)
}
