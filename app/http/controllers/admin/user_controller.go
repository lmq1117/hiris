package admin

import (
	"github.com/kataras/iris"
	"strconv"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetBy() {
	//this.Ctx.JSON(iris.Map{"boy": "goods jobs"})
	c.Ctx.HTML("hello! your id is : <h1>" + strconv.FormatInt(1, 10) + "</h1>")
}
