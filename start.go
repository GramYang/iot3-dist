package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression),
		static.Serve("/", static.LocalFile("dist", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	_ = r.Run(":8085")
}
