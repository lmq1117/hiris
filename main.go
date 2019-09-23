package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/mvc", func(ctx iris.Context) {
		ctx.HTML("<b>welcome to mvc hello world!</b>")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Handle("GET", "hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "hello iris mvc"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
