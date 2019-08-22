package admin

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/mvc"
	"hiris/app/http/models"
	//"strconv"
)

var (
	user models.User
)
var db = models.ConnectDB()

type UserController struct {
	Ctx iris.Context
}

type JsonResult struct {
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

// admin/user/haha
func (c *UserController) GetHaha() string {
	//c.Ctx.ReadJSON()
	user.ID = 6
	db.Find(&user)
	return user.Name
	//return "====" +strconv.Itoa(user.ID) + "====" + "===="
}
