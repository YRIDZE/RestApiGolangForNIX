package main

import (
	"github.com/labstack/echo/v4"
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

func DeleteComments(c echo.Context) error {
	postId, _ := strconv.Atoi(c.Param("postId"))
	for index, item := range comms {
		if item.PostId == postId {
			comms = append(comms[:index], comms[index+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, comms)
}

func PutComments(c echo.Context) error {
	postId, _ := strconv.Atoi(c.Param("postId"))
	for index, item := range comms {
		if item.PostId == postId {
			comms = append(comms[:index], comms[index+1:]...)
			var comm Comments
			c.Bind(&comm)
			comm.PostId = postId
			comms = append(comms, comm)
			//return c.JSON(http.StatusOK, comm)
		}
	}
	return c.JSON(http.StatusOK, comms)
}

func PostComments(c echo.Context) error {
	var comm Comments
	c.Bind(&comm)
	comms = append(comms, comm)
	return c.JSON(http.StatusOK, comms)
}

func GetComment(c echo.Context) error {
	postId, _ := strconv.Atoi(c.Param("postId"))
	for _, item := range comms {
		if item.PostId == postId {
			return c.JSON(http.StatusOK, item)
		}

	}
	return c.JSON(http.StatusOK, comms)

}

func GetComments(c echo.Context) error {
	return c.JSON(http.StatusOK, comms)
}

func Setup() *echo.Echo {
	echo := echo.New()
	e := echo.Group("/comms")
	{
		e.GET("", GetComments)
		e.GET("/:postId", GetComment)
		e.POST("", PostComments)
		e.PUT("/:postId", PutComments)
		e.DELETE("/:postId", DeleteComments)
	}
	return echo
}

func main() {
	comms = append(comms, Comments{1, 2, "OOOOne", "Two", "Three"})

	e := Setup()
	e.Logger.Fatal(e.Start(":8000"))

}
