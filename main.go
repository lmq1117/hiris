package main

import "github.com/kataras/iris"

type User struct {
	Username string `json:"username"`
	City     string `json:"city"`
	Age      int    `json:"age"`
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	//为特殊http错误自定义错误返回
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("服务器内部错误：%s", errMessage)
			return
		}
		ctx.Writef("未知服务器错误")
	})

	//GET
	app.Get("/encode", func(ctx iris.Context) {
		doe := User{
			Username: "王大锤1",
			City:     "西部小山村",
			Age:      25,
		}
		ctx.JSON(doe)
	})

	//Run
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}
