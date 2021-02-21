package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserApi struct{}

type Comments struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var comms []Comments

func (u *UserApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodPut:
		doPut(w, r)
	case http.MethodDelete:
		doDelete(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
	}

}

func doPost(w http.ResponseWriter, r *http.Request) {
	var comm Comments
	_ = json.NewDecoder(r.Body).Decode(&comm)
	comms = append(comms, comm)
	json.NewEncoder(w).Encode(comms)
}

func doGet(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comms)
}

func doPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, _ := strconv.Atoi(r.URL.Query().Get("postId"))
	fmt.Println(params)
	for index, item := range comms {
		if item.PostId == params {
			comms = append(comms[:index], comms[index+1:]...)
			var comm Comments
			_ = json.NewDecoder(r.Body).Decode(&comm)
			comm.PostId = params
			comms = append(comms, comm)
			json.NewEncoder(w).Encode(comm)
		}
	}
	json.NewEncoder(w).Encode(comms)
}

func doDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, _ := strconv.Atoi(r.URL.Query().Get("postId"))
	for index, item := range comms {
		if item.PostId == params {
			comms = append(comms[:index], comms[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(comms)
}

func main() {
	comms = append(comms, Comments{1, 2, "One", "Two", "Three"})

	http.Handle("/comms", new(UserApi))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
