package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/open-door", func(c *gin.Context) {
        OpenDoor()
        c.JSON(200, gin.H{
            "status": "open",
        })
    })
    r.Run(":80")
}
