package _type

import (
	"fmt"
	"log"
	"testing"
)

func TestArticleGenerator(t *testing.T) {
	a,err := ArticleGenerator("../static/1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

func TestSyncArticle(t *testing.T) {
	SyncArticle("../static/")
	log.Println(IDToArt["0"].Year)
}
