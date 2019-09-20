package main

import (
	"github.com/kataras/iris"
	"net/http"
	"strings"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>资源未找到</b>")
	})

	//todo ServerFile 干嘛的呀
	app.Get("/abc", func(ctx iris.Context) {
		ctx.ServeFile("./public/index.html", false)
	})
	app.Get("/profile/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("username"))
	})

	app.HandleDir("/", "./public")

	myOtherHandler := func(ctx iris.Context) {
		ctx.Writef("在我们的自定义路由器包装器手动触发的处理程序内部")
	}

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, route http.HandlerFunc) {
		path := r.URL.Path
		if strings.HasPrefix(path, "/other") {
			ctx := app.ContextPool.Acquire(w, r)
			myOtherHandler(ctx)
			app.ContextPool.Release(ctx)
			return
		}
		route.ServeHTTP(w, r)
	})
	return app
}
