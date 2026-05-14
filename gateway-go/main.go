package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "healthy",
            "engine": "woogle-gateway",
        })
    })

    r.GET("/api/search", func(c *gin.Context) {
        query := c.Query("q")

        c.JSON(http.StatusOK, gin.H{
            "query": query,
            "results": []string{
                "Recursive sovereign search online",
                "Distributed execution node",
                "Semantic memory vector result",
            },
        })
    })

    r.Run(":8080")
}