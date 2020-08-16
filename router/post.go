package router

import (
	"github.com/gin-gonic/gin"
	_type "github.com/mangoqiqi/Blog-Backend/type"
	"net/http"
	"strings"
)

// Post xxx
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Year    string `json:"year"`
	Summary string `json:"summary"`
}

// GetBlogsByYea func return
// @GET /api/blogs/year/2020
func GetBlogsByYear(c *gin.Context) {
	var res []Post
	y := c.Param("year")
	for _, v := range _type.IDToArt {
		if strings.Contains(v.Year, y){
			res = append(res, Post{
				ID:      v.ID,
				Title:   v.Title,
				Date:    v.Date,
				Year:    v.Year,
				Summary: v.Summary,
			})
		}
	}
	c.JSON(http.StatusOK, res)
}

// @GET /api/blogs/kind/docker
func GetBlogsByKind(c *gin.Context) {
	var res []Post
	k := c.Param("kind")
	for _, v := range _type.KindToID[k] {
		art := _type.IDToArt[v]
		res = append(res, Post{
			ID:      art.ID,
			Title:   art.Title,
			Date:    art.Date,
			Year:    art.Year,
			Summary: art.Summary,
		})
	}
	c.JSON(http.StatusOK, res)
}

// GetBlogDetail func return
// @GET /api/detail/2
func GetBlogDetail(c *gin.Context) {
	type postDetail struct {
		ID      string    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	type respon struct {
		Result struct {
			Code    int
			meesage string
		} `json:"result"`
		Detail postDetail `json:"detail"`
	}
	ID := c.Param("id")
	res := respon{}
	if _,ok := _type.IDToArt[ID]; !ok {
		res.Result.Code = -1
		res.Result.meesage = "filed"
	} else {
		res.Result.Code = 0
		res.Result.meesage = "success"
		art := _type.IDToArt[ID]
		res.Detail = postDetail{
			ID:      art.ID,
			Title:   art.Title,
			Content: art.Content,
		}
	}
	c.JSON(http.StatusOK, res)
}

// GetAllBlogs func return
// @GET /api/blogs/all?filter=key
func GetAllBlogs(c *gin.Context) {
	filterBy := c.DefaultQuery("filterBy", "")
	var res []Post
	if filterBy != "" {
		for _, v := range _type.IDToArt {
			if strings.Contains(v.Title, filterBy) ||
				strings.Contains(v.Summary, filterBy) ||
				strings.Contains(v.Content, filterBy) {
				res = append(res, Post{
					ID:      v.ID,
					Title:   v.Title,
					Date:    v.Date,
					Year:    v.Year,
					Summary: v.Summary,
				})
			}
		}
	} else {
		for _, v := range _type.IDToArt {
			res = append(res, Post{
				ID:      v.ID,
				Title:   v.Title,
				Date:    v.Date,
				Year:    v.Year,
				Summary: v.Summary,
			})
		}
	}
	c.JSON(http.StatusOK, res)
}
