package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"hiris/app/http/models"
)

var users []models.User
var user models.User

func init() {
	initDB() //初始化db
}
func main() {

	//insert
	//user := models.User{Name: "司马昭", Age: 35, Birthday: time.Now()}
	//fmt.Println(db.NewRecord(user))
	//db.Create(&user)
	//fmt.Println(db.NewRecord(user))

	//user := models.User{}
	//db.First(&user)
	//db.Last(&user)
	//fmt.Println(user)

	//users 是 切片 类型是models.User
	//db.Find(&users)
	//fmt.Println(users)
	//fmt.Println(users[1].CreatedAt)

	//按主键获取
	//db.First(&user, 10)

	//简单sql
	//db.Find(&user, "name = ?", "孔明")
	//db.Find(&users, "name <> ? AND age > ?", "黄忠", 55)

	//Struct
	//db.Find(&users, models.User{Age: 55})
	//db.Find(&users, map[string]interface{}{"age": 60})

	//================= Not 条件
	// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;
	//db.Not("name","刘备").First(&user)

	//SELECT * FROM user WHERE name NOT IN ("曹操", "曹丕");
	//db.Not("name",[]string{"曹操","曹丕"}).Find(&users)

	//SELECT * FROM users WHERE id NOT IN (1,2,3);
	//db.Not([]int64{1,2,3}).First(&user)

	//SELECT * FROM users limit 1
	//db.Not([]int64{}).First(&user)

	//Plain sql
	//db.Not("name = ?","孔明").First(&user)

	//struct
	//db.Not(models.User{Name:"刘备"}).First(&user)

	//===========带内联条件的查询

	//fmt.Println(user)
	//fmt.Println(users)

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello iris</b>")
	})
	//config := iris.TOML("./config/iris.tml")
	//fmt.Println(config.Other["MyServerName"])
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./config/iris.tml")))

}

func initDB() {
	//连接数据库
	models.ConnectDB()
}
