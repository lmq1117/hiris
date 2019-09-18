package main

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"hiris/app"
	//"hiris/app/http/models"
	"github.com/kataras/iris"
	"regexp"
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

	//routing.api
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

	//routing.offlineRoute
	//None 从外部隐藏路由，可以通过Context.Exec方法从handle调用
	none := app.None("/invisible/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s with method:%s", ctx.Params().Get("username"), ctx.Method())
		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/change", func(ctx iris.Context) {
		if none.IsOnline() {
			none.Method = iris.MethodNone
		} else {
			none.Method = iris.MethodGet
		}
		app.RefreshRouter()
	})

	app.Get("/execute", func(ctx iris.Context) {
		if !none.IsOnline() {
			ctx.Values().Set("from", "/execute with offline access")
			ctx.Exec("NONE", "/invisible/iris")
			return
		}
		ctx.Values().Set("from", "/execute")
		ctx.Exec("GET", "/invisible/iris")
	})

	//路由分组方式一 groupingRoutes
	users := app.Party("/users", myMiddleware)
	{
		users.Get("/{id:uint64}/profile", func(ctx iris.Context) {
			ctx.Writef("<h1>Hello from /users/%s/profile </h1>", ctx.Params().Get("id"))
		})
		users.Get("/message/{id:uint64}", func(ctx iris.Context) {
			ctx.Writef("<h1>Hello from /users/message/%s </h1>", ctx.Params().Get("id"))
		})
	}

	//路由分组方式二 PartyFunc
	app.PartyFunc("/partyfunc", func(p iris.Party) {
		p.Use(myMiddleware)
		p.Get("/{id:uint64}/profile", func(ctx iris.Context) {
			ctx.Writef("<h1>Hello from /partyfunc/%s/profile </h1>", ctx.Params().Get("id"))
		})
		p.Get("/message/{id:uint64}", func(ctx iris.Context) {
			ctx.Writef("<h1>Hello from /partyfunc/message/%s </h1>", ctx.Params().Get("id"))
		})
	})

	//区分路由路径结尾是否带 /
	// http://localhost:8080/user/1/ Result:Not Found
	// http://localhost:8080/user/1 Result:User ID: 1
	//app.Run(iris.Addr(":8080"),iris.WithoutPathCorrection)

	/*
		|--------------------------------------------------------------------------
		| Route path paramter types
		|--------------------------------------------------------------------------
		|
		| 1. path
		| 2. string
		| 3. int min(5)
		|
	*/
	//匹配 /assets/*/* eg /assets/aaa/bbb/ccc
	app.Get("/assets/{asset:path}", func(ctx iris.Context) {
		ctx.Writef(ctx.Path() + "|参数值:" + ctx.Params().GetString("asset") + "\n" + ctx.Params().Get("asset"))
	})

	app.Get("/profile/me", func(ctx iris.Context) {
		ctx.Writef(ctx.Path())
	})

	app.Get("/profile/{username:string}", func(ctx iris.Context) {
		ctx.Writef(ctx.Path() + "|参数值:" + ctx.Params().GetString("username"))
	})

	app.Get("/u/{userid:int min(5)}", func(ctx iris.Context) {
		ctx.Writef(ctx.Path() + "|参数值:" + ctx.Params().GetString("userid"))
	})

	app.Get("alp/{name:alphabetical max(5)}", func(ctx iris.Context) {
		ctx.Writef(ctx.Path() + "|参数值:" + ctx.Params().Get("name"))
	})
	//自定义路由参数验证规则
	latLonExpr := "^-?[0-9]{1,2}(?:\\.[0-9]{1,4})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)
	app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}", func(ctx iris.Context) {
		ctx.Writef("Lat:%s|Lon:%s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
	})

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
