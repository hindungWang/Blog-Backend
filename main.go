package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.Use(cors.Default())
	// @GET /api/blogs/year/2020
	r.GET("/api/blogs/year/:year", func(c *gin.Context) {
		type respon struct {
			ID 		int     `json:"id"`
			Title 	string  `json:"title"`
			Date 	string  `json:"date"`
			Year 	string  `json:"year"`
			Summary string  `json:"summary"`
		}
		var res []respon
		_ = c.Param("year")
		for i:=1;i<20;i++ {
			one := respon{
				ID: i,
				Title: "ttt1",
				Date: "fsd1",
				Year: "rfse1",
				Summary: `ksdjfaj
				dsfsjsdfsf***fgdsfg***`,
			}
			res = append(res, one)
		}
        c.JSON(http.StatusOK, res)
	})
	
	r.GET("/api/detail/:id", func(c *gin.Context)  {
		type respon struct {
			ID 		int     `json:"id"`
			Title 	string  `json:"title"`
			Content string  `json:"content"`
		}
		ID := c.Param("id")
		id,_ := strconv.Atoi(ID)
		res := respon{
			ID: id,
			Title: "ddasdfsadf",
			Content: `凡萨达号房送到附近萨达。
			***海上的凡***，sufism发射的防护发生较大和方式的;
			大概的法官的凡`,
		}
		c.JSON(http.StatusOK, res)
	})
    r.Run()
}