
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct { //带结构标签，反引号来包围字符串
	Id int  `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comment []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("json/post.json")
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err!= nil {
		fmt.Println("error reading json file")
		return
	}
	var post Post
	json.Unmarshal(jsonData,&post)
	fmt.Println(post)
}
