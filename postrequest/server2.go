package main

import (
	"net/http"

	"github.com/labstack/echo"
)

//User 構造体
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/save", save)
	e.Logger.Fatal(e.Start(":8080"))
}

func save(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
