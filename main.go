package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hindungWang/Blog-Backend/router"
	_type "github.com/hindungWang/Blog-Backend/type"
	"log"
)

func Sync(p string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				go _type.SyncArticle(p)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(p)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

const Path = "./static/"
func main() {
	r := gin.Default()
	r.Use(cors.Default())

	_type.SyncArticle(Path)
	r.GET("/api/blogs/year/:year", router.GetBlogsByYear)
	r.GET("/api/blogs/kind/:kind", router.GetBlogsByKind)
	r.GET("/api/detail/:id", router.GetBlogDetail)
	r.GET("/api/blogs/all", router.GetAllBlogs)
	go Sync(Path)
	err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}