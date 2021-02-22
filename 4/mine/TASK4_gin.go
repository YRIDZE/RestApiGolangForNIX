package main

import (
	"github.com/gin-gonic/gin"
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

var comms []Comments

func GetComments(context *gin.Context) {
	context.JSON(http.StatusOK, comms)
}
func GetComment(context *gin.Context) {
	postId, _ := strconv.Atoi(context.Params.ByName("postId"))
	for _, item := range comms {
		if item.PostId == postId {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, comms)

}
func PostComments(context *gin.Context) {
	var comm Comments
	context.BindJSON(&comm)
	comms = append(comms, comm)
	context.JSON(http.StatusOK, comm)

}

func PutComments(context *gin.Context) {
	postId, _ := strconv.Atoi(context.Params.ByName("postId"))
	for index, item := range comms {
		if item.PostId == postId {
			comms = append(comms[:index], comms[index+1:]...)
			var comm Comments
			context.BindJSON(&comm)
			comm.PostId = postId
			comms = append(comms, comm)
			context.JSON(http.StatusOK, comm)
		}
	}
	context.JSON(http.StatusOK, comms)

}

func DeleteComments(context *gin.Context) {
	postId, _ := strconv.Atoi(context.Params.ByName("postId"))
	for index, item := range comms {
		if item.PostId == postId {
			comms = append(comms[:index], comms[index+1:]...)
			break
		}
	}
	context.JSON(http.StatusOK, comms)

}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	r := router.Group("/comms")
	{
		r.GET("", GetComments)
		r.GET("/:postId", GetComment)
		r.POST("", PostComments)
		r.PUT("/:postId", PutComments)
		r.DELETE("/:postId", DeleteComments)

	}
	return router
}

func main() {
	comms = append(comms, Comments{1, 2, "OOOOne", "Two", "Three"})

	router := SetupRouter()
	router.Run(":8000")
}
