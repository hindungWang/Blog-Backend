package _type

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

type Test struct {
	ID      string         `mapstructure:"Id"`
	Title   string         `mapstructure:"Title"`
	Date    string         `mapstructure:"Date"`
	Year    string         `mapstructure:"Year"`
	Summary string         `mapstructure:"Summary"`
	Tags    []string       `mapstructure:"Tags"`
	Content string         
}

func TestVisper(t *testing.T) {

	fd, err := os.Open("../static/0.md")
	if err != nil {
		t.Fatal(err.Error())
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

	fmt.Println(string(config))
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
	ts := &Test{}
	viper.Unmarshal(ts)

	fmt.Println(ts)
}
