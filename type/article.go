package _type

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const filePath = "../static/"

// IDToArt ID --> article
var IDToArt map[string]*Article
// KindToID xxx
var KindToID map[string][]string

type Article struct {
	ID      string   `mapstructure:"Id"`
	Title   string   `mapstructure:"Title"`
	Date    string   `mapstructure:"Date"`
	Year    string   `mapstructure:"Year"`
	Summary string   `mapstructure:"Summary"`
	Tags    []string `mapstructure:"Tags"`
	Content string
}

// ArticleGenerator return
func ArticleGenerator(path string) (*Article, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	var flag int
	var config []byte
	var content string
	sp := []byte("\n")

	for {
		res, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		if flag < 2 {
			if strings.Contains(string(res), "---") {
				flag++
				continue
			}
			config = append(config, res...)
			config = append(config, sp...)
		} else {
			content += fmt.Sprintf("%s\n", string(res))
		}
	}

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	article := &Article{}

	err = viper.Unmarshal(article)
	if err != nil {
		return nil, err
	}
	article.Content = content
	return article, nil
}

// SyncArticle return
func SyncArticle(pathDir string) {
	Art := make(map[string]*Article)
	Index := make(map[string][]string)
	files, _ := ioutil.ReadDir(pathDir)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			art, err := ArticleGenerator(pathDir + file.Name())
			if err != nil {
				log.Println(err.Error())
				continue
			}
			Art[art.ID] = art
			for _,v := range art.Tags {
				tmp := Index[v]
				tmp = append(tmp, art.ID)
				Index[v] = tmp
			}
		}
	}
	IDToArt = Art
	KindToID = Index
	//log.Println(IDToArt)
	//log.Println(KindToID)
}
