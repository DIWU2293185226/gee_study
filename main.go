package main

import (
	"go_study/gee/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCodes, c.Rep.RequestURI, time.Since(t))
	}
}

func main() {
	//day2
	// r := gee.NewEngine()
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// })
	// r.GET("/hello", func(c *gee.Context) {
	// 	// expect /hello?name=geektutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })

	// r.POST("/login", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })

	// r.Run(":9999")

	//day3
	r := gee.NewEngine()
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// })

	// r.GET("/hello", func(c *gee.Context) {
	// 	// expect /hello?name=geektutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })

	// r.GET("/hello/:name", func(c *gee.Context) {
	// 	// expect /hello/geektutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// })

	// r.GET("/assets/*filepath", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	// })

	//分组控制测试
	// r.GET("/index", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// })
	// v1 := r.Group("/v1")
	// {
	// 	v1.Get("/", func(c *gee.Context) {
	// 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// 	})

	// 	v1.Get("/hello", func(c *gee.Context) {
	// 		// expect /hello?name=geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// 	})
	// }
	// v2 := r.Group("/v2")
	// {
	// 	v2.Get("/hello/:name", func(c *gee.Context) {
	// 		// expect /hello/geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// 	})
	// 	v2.Post("/login", func(c *gee.Context) {
	// 		c.JSON(http.StatusOK, gee.H{
	// 			"username": c.PostForm("username"),
	// 			"password": c.PostForm("password"),
	// 		})
	// 	})

	// }
	/*
		curl "http://localhost:9999/v1/hello?name=geektutu"
		curl "http://localhost:9999/v2/hello/geektutu"
	*/

	//中间件测试
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.Get("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
