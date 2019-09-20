package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"reflect"
)

type MyContext struct {
	iris.Context
}

var _ iris.Context = &MyContext{}

func (ctx *MyContext) Do(handler context.Handlers) {
	context.Do(ctx, handler)
}

func (ctx *MyContext) Next() {
	context.Next(ctx)
}

func (ctx *MyContext) HTML(format string, args ...interface{}) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from MyContext")
	ctx.ContentType("text/html")
	return ctx.Writef(format, args...)
}

func main() {
	app := iris.New()
	app.ContextPool.Attach(func() iris.Context {
		return &MyContext{
			Context: context.NewContext(app),
		}
	})
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Handle("GET", "/", recordWhichContextForExample, func(ctx iris.Context) {
		ctx.HTML("<h1>Hello from my custom context 's HTML! </h1>")
	})
	app.Handle("GET", "/hi/{name:alphabetical}", recordWhichContextForExample, func(ctx iris.Context) {
		name := ctx.Params().GetString("name")
		ctx.ViewData("name", name)
		ctx.Gzip(true)
		ctx.View("hi.html")
	})
	app.Run(iris.Addr(":8080"))

}

func recordWhichContextForExample(ctx iris.Context) {
	ctx.Application().Logger().Infof("(%s) Handler is executing from:'%s'", ctx.Path(), reflect.TypeOf(ctx).Elem().Name())
	ctx.Next()
}
