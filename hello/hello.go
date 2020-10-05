package hello

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	user2 struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user2{}
	seq   = 1
)

func getUser2(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}
