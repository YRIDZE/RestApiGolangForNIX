package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Comments struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Posts struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func f(url string, db *gorm.DB) {
	const a, b, c = 3, 4, "foo"
	resp, err := http.Get(url + "/posts?userId=7")
	if err != nil {
		fmt.Println("error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var posts []Posts
	err = json.Unmarshal(body, &posts)

	for i := 0; i < len(posts); i++ {
		db.Create(&posts)
		if err != nil {
			panic(err)
		}
		go f1(url, posts[i].Id, db)
	}
}

func f1(url string, index int, db *gorm.DB) {
	resp, err := http.Get(url + "/comments?postId=" + strconv.Itoa(index))
	if err != nil {
		fmt.Println("error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var comms []Comments
	err = json.Unmarshal(body, &comms)

	for i := 0; i < len(comms); i++ {
		db.Create(&comms)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := gorm.Open(mysql.Open("root:08101999@(localhost:3306)/post_new"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Posts{})
	db.AutoMigrate(&Comments{})
	url := "https://jsonplaceholder.typicode.com"

	go f(url, db)
	var input string
	_, _ = fmt.Scanln(&input)

}
