package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getting(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func posting(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func putting(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleting(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func patching(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func head(c *gin.Context) {
  c.Status(http.StatusOK)
}

func options(c *gin.Context) {
  c.Status(http.StatusOK)
}

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  router.GET("/someGet", getting)
  router.POST("/somePost", posting)
  router.PUT("/somePut", putting)
  router.DELETE("/someDelete", deleting)
  router.PATCH("/somePatch", patching)
  router.HEAD("/someHead", head)
  router.OPTIONS("/someOptions", options)
  
  router.Run() // listens on 0.0.0.0:8080 by default
}