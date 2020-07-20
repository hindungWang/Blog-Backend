package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mangoqiqi/Blog-Backend/router"
	"log"
)

func main()  {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/blogs/year/:year", router.GetBlogsByYear)
	r.GET("/api/blogs/kind/:kind", router.GetBlogsByKind)
	
	r.GET("/api/detail/:id", router.GetBlogDetail)
    err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}