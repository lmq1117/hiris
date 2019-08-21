package app

import (
	"hiris/app/http/controllers/admin"
	"hiris/app/http/models"
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
)

func InitIris() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.OnAnyErrorCode(func(ctx context.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "admin123") {
			_, err = ctx.JSON(iris.Map{strconv.Itoa(ctx.GetStatusCode()): "http on any error"}) //todo 可以更详细些
		}
		if err != nil {
			logrus.Error(err)
		}
	})

	mvc.Configure(app.Party("/admin"), func(m *mvc.Application) {
		//m.Router.Use(middleware.AdminAuth) 中间件 //todo
		m.Router.Use(func(ctx iris.Context) {
			ctx.Application().Logger().Infof("Path:%s", ctx.Path())
			ctx.Next()
		})
		m.Party("/user").Handle(new(admin.UserController))
	})
	server := &http.Server{Addr: ":8088"}
	handleSignal(server)
	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		logrus.Infof("got signal [%s],exiting now", s)
		if err := server.Close(); err != nil {
			logrus.Infof("server close failed: " + err.Error())
		}
		models.CloseDB()
		logrus.Infof("Exited")
		os.Exit(0)
	}()
}
