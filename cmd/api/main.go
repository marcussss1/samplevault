package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: false,
	}))

	//conn, err := tarantool.Connect("tarantool-0.tarantool.default.svc.cluster.local:3301", tarantool.Opts{
	//	User: "tarantool",
	//	Pass: "tarantool",
	//})
	//if err != nil {
	//	fmt.Println("Connection refused")
	//	return
	//}
	//defer conn.Close()

	e.GET("/api/v1", func(ctx echo.Context) error {
		//resp, err := conn.Select("tarantool", "primary", 0, 1, tarantool.IterEq, []interface{}{3})
		//if err != nil {
		//	fmt.Println(err)
		//	return err
		//}
		//
		//fmt.Println(resp)
		var a []int
		fmt.Println(a[0])

		fmt.Println("VERSION 2 REQUEST")

		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "VERSION 2 HANDLE",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
