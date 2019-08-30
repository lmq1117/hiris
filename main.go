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

	//区分路由路径结尾是否带 /
	// http://localhost:8080/user/1/ Result:Not Found
	// http://localhost:8080/user/1 Result:User ID: 1
	//app.Run(iris.Addr(":8080"),iris.WithoutPathCorrection)

	//不区分路由路径结尾是否带 /
	app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection)
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
