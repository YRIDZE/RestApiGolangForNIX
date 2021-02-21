package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

/*type Comments struct {
	PostId string `json:"postId"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}*/

type Posts struct {
	UserId string `json:"userId"`
	Id     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var db *gorm.DB
var posts []Posts

//var comms []Comments

func Connect() {
	d, err := gorm.Open(mysql.Open("root:08101999@(localhost:3306)/post_new"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}

func createPosts(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var post Posts
	_ = json.NewDecoder(req.Body).Decode(&post)
	post.Id = params["id"]
	posts = append(posts, post)
	json.NewEncoder(w).Encode(posts)
}
func getPosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func getPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range posts {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Posts{})

}
func updatePosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range posts {
		if item.Id == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Posts
			_ = json.NewDecoder(req.Body).Decode(&post)
			post.Id = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
		}
	}
	json.NewEncoder(w).Encode(posts)

}
func deletePosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range posts {
		if item.Id == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func main() {
	/*Connect()
	db:= GetDB()
	db.AutoMigrate(&Posts{})
	db.AutoMigrate(&Comments{})*/
	router := mux.NewRouter()

	posts = append(posts, Posts{"12", "11", "HelloHell", "HoHoHo"})
	posts = append(posts, Posts{"15", "16", "ByeByyye", "HeHeHe"})

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", createPosts).Methods("POST")
	router.HandleFunc("/posts/{id}", updatePosts).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePosts).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
