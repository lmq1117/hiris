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

	userRoute := app.Party("/user", logThisMiddleware)

	//花括号 跟 上边这句 没啥关系
	{
		userRoute.Get("/{id:int min(1)}", getUserByID)
		userRoute.Get("/{name:string}/{city:string}", getUserByName)
	}

	//Run
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("====Path:%s | IP: %s====", ctx.Path(), ctx.RemoteAddr()) //格式化LOG
	ctx.Next()
}

func getUserByID(ctx iris.Context) {
	userId := ctx.Params().Get("id")
	user := User{Username: "username" + userId}
	ctx.XML(user)
}

func getUserByName(ctx iris.Context) {
	name := ctx.Params().Get("name")
	city := ctx.Params().Get("city")
	//age := ctx.Params().Values().GetInt("age")
	user := User{Username: name, City: city}
	ctx.JSON(user)
}
