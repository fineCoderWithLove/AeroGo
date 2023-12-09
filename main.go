package main

import (
	"aerogo"
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

	r.GET("/hello/:name", func(c *aerogo.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *aerogo.Context) {
		c.JSON(http.StatusOK, aerogo.H{"filepath": c.Param("filepath")})
	})

	r.Run(":3333")

}
