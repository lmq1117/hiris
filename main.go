package main

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"hiris/app"
	//"hiris/app/http/models"
	"github.com/kataras/iris"
)

//var users []models.User
//var user models.User

//func init() {
//	initDB() //初始化db
//}
//func main() {
//	app.InitIris()
//	initDB()
//}
//
//func initDB() {
//	//连接数据库
//	models.ConnectDB()
//}

func main() {
	app := iris.New()
	app.Use(myMiddleware)
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello iris World!")
		ctx.View("hello.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Handle("GET", "/contact", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello from /contact </h1>")
	})

	app.Get("/method", handler)
	app.Post("/method", handler)
	app.Put("/method", handler)
	app.Delete("/method", handler)
	app.Options("/method", handler)
	app.Trace("/method", handler)
	app.Connect("/method", handler)
	app.Head("/method", handler)
	app.Patch("/method", handler)
	app.Any("/methodany", handler)

	//区分路由路径结尾是否带 /
	// http://localhost:8080/user/1/ Result:Not Found
	// http://localhost:8080/user/1 Result:User ID: 1
	//app.Run(iris.Addr(":8080"),iris.WithoutPathCorrection)

	//不区分路由路径结尾是否带 /
	app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection)
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s, method %s\n", ctx.Path(), ctx.Method())
	ctx.Next()
}

func handler(ctx iris.Context) {
	ctx.Writef("Hello from method %s and path %s\n", ctx.Method(), ctx.Path())
}
