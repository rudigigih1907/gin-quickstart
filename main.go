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

  // This handler will match /user/john but will not match /user/ or /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // However, this one will match /user/john/ and also /user/john/send
  // If no other routers match /user/john, it will redirect to /user/john/
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  // Query string parameters are parsed using the existing underlying request object.
  // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
  router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })
  
  router.Run() // listens on 0.0.0.0:8080 by default
}