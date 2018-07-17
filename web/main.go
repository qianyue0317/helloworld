package main

import "github.com/astaxie/beego"

/**
 * beego 演示demo
 */
type HomeController struct {
	beego.Controller
}

func (controller *HomeController) Get()  {
	controller.Ctx.WriteString("hello world")
}



func main() {
	beego.Router("/",&HomeController{})
	beego.Run()
}
