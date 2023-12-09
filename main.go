package main

import (
	"aerogo"
	"log"
	"net/http"
)

func main() {
	r := aerogo.New()
	r.GET("/", func(c *aerogo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *aerogo.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *aerogo.Context) {
		log.Println(c.Req)
		c.JSON(http.StatusOK, aerogo.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
