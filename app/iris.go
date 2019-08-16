package app

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/sirupsen/logrus"
	"hiris/app/http/controllers/admin"
)

func InitIris() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.OnAnyErrorCode(func(ctx context.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "admin") {
			_, err = ctx.JSON(iris.Map{strconv.Itoa(ctx.GetStatusCode()): "http on any error"}) //todo 可以更详细些
		}
		if err != nil {
			logrus.Error(err)
		}
	})

	mvc.Configure(app.Party("/admin"), func(m *mvc.Application) {
		//m.Router.Use(middleware.AdminAuth) 中间件 //todo
		m.Party("/user").Handle(new(admin.UserController))
	})
	server := &http.Server{Addr: ":8080"}
	handleSignal(server)
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		logrus.Infof("got signal [%s],exiting now", s)
	}()
}
