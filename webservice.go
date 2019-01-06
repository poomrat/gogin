package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	captcha "github.com/poomrat/gocaptcha"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/captcha", func(c *gin.Context) {
		rand.Seed(time.Now().UTC().UnixNano())
		c.String(200, captcha.Generate(rand.Intn(2)+1, rand.Intn(10), rand.Intn(3)+1, rand.Intn(10)))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
