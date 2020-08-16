package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/bmonikraj/goql/routes"
)

var router *gin.Engine

func main() {

    router = gin.Default()

    router = routes.InitializeRoutes(router)

    host := os.Args[1]
    port := os.Args[2]
    router.Run(host+":"+port)
}