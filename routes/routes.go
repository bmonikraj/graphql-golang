package routes

import (
    "github.com/gin-gonic/gin"
	"github.com/bmonikraj/goql/handler"
)

func InitializeRoutes(router *gin.Engine) *gin.Engine {
	router.POST("/graphql", handler.GraphQLHandler)
	return router
}