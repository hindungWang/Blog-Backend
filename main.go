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

	r.GET("/api/blogs/year/:year", router.GetBlogsByYea)
	
	r.GET("/api/detail/:id", router.GetBlogDetail)
    err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}