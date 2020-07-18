package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Post xxx
type Post struct {
	ID 		int     `json:"id"`
	Title 	string  `json:"title"`
	Date 	string  `json:"date"`
	Year 	string  `json:"year"`
	Summary string  `json:"summary"`
}


// GetBlogsByYea func return
// @GET /api/blogs/year/2020
func GetBlogsByYea(c *gin.Context) {
	var res []Post
	_ = c.Param("year")
	for i:=1;i<20;i++ {
		one := Post{
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
}

// GetBlogDetail func return
// @GET //api/detail/2
func GetBlogDetail(c *gin.Context)  {
	type postDetail struct {
		ID 		int     `json:"id"`
		Title 	string  `json:"title"`
		Content string  `json:"content"`
	}
	type respon struct {
		Result struct {
			Code int
			meesage string
		} `json:"result"`
		Detail postDetail`json:"detail"`
	}
	ID := c.Param("id")
	id,_ := strconv.Atoi(ID)
	res := respon{}
	res.Result.Code = 0
	res.Result.meesage = "success"
	fi, err := os.Open("./1.md")
    if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
    }
    defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	log.Println(string(fd))
	res.Detail = postDetail{
		ID: id,
		Title: "ddasdfsadf",
		Content: string(fd),
	}
	c.JSON(http.StatusOK, res)
}