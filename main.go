package main

import (
	"gee/gee"
	"net/http"
)

//type Engine struct {
//}
//
//func (receiver *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
//
//	}
//}

func main() {
	r := gee.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v2.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	//r.POST("/login", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})
	r.Run("9999")
}
