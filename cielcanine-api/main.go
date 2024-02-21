package main


import (
    "github.com/gin-gonic/gin"
    "myproject"
)

func main() {
    router := gin.Default()
    myproject.SetupRoutes(router)
    router.Run(":8080")
}