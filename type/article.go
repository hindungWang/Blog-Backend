package _type

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const filePath = "../static/"

var IDToArt map[string]*Article

type Article struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Date    string         `json:"date"`
	Year    string         `json:"year"`
	Summary string         `json:"summary"`
	Tags    map[string]int `json:"tags"`
	Content string         `json:"content"`
}

func ArticleGenerator(path string) (*Article, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	var line int
	var str []string
	var content string
	for {
		a, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		if line < 5 {
			str = append(str, string(a))
		} else {
			content += fmt.Sprintf("%s\n", string(a))
		}
		line++
	}
	if len(str) < 5 {
		return nil, fmt.Errorf("read err!")
	}
	var title string = str[0][6:]
	var date string = str[1][5:]
	var year string = str[2][5:]
	var sum string = str[3][8:]
	var tag []string = strings.Split(str[4][4:], ",")
	tags := make(map[string]int)
	for _,v := range tag {
		tags[v]++
	}
	article := &Article{
		ID:      getIDFromPath(path),
		Title:   title,
		Date:    date,
		Year:    year,
		Summary: sum,
		Tags:    tags,
		Content: content,
	}
	return article, nil
}

func getIDFromPath(p string) string {
	i := len(p) - 1
	for ; i >= 0; i-- {
		if p[i] == '/' {
			break
		}
	}
	var res string
	for j := i + 1; j < len(p); j++ {
		if p[j] == '.' {
			break
		}
		res += string(p[j])
	}
	return res
}

func SyncArticle(pathDir string) {
	Art := make(map[string]*Article)
	files, _ := ioutil.ReadDir(pathDir)
	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			art, err := ArticleGenerator(pathDir + file.Name())
			if err != nil {
				log.Println(err.Error())
				continue
			}
			Art[art.ID] = art
		}
	}
	IDToArt = Art
}
