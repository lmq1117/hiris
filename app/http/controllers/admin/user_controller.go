package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"hiris/app/http/models"
	//"github.com/kataras/iris/mvc"
	//"strconv"
)

var (
	db   *gorm.DB
	user models.User
)

type UserController struct {
	Ctx iris.Context
}

// admin/user/id
func (c *UserController) GetId() int64 {
	//this.Ctx.JSON(iris.Map{"boy": "goods jobs"})
	//c.Ctx.HTML("hello! your id is : <h1>" + strconv.FormatInt(1, 10) + "</h1>")
	return 1
}

//admin/user/userinfo
func (c *UserController) GetUserinfo() string {
	return "user info method"
}

func (c *UserController) GetHaha() string {
	db.First(&user)
	return user.Name
}
