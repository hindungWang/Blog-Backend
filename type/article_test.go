package _type

import (
	"fmt"
	"testing"
)

func TestArticleGenerator(t *testing.T) {
	a,_ := ArticleGenerator("../static/0.md")
	fmt.Println(a)
}

func TestSyncArticle(t *testing.T) {
	SyncArticle("../static/")
}
