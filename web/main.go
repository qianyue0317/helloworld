package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id  int
	Pre string
	Ng  string
	Ser string
	Th  string `orm:"size(500)"`
}

/**
 * beego 演示demo
 */
type HomeController struct {
	beego.Controller
}



func (controller *HomeController) Get() {
	controller.Ctx.WriteString("hello world")
}

func main() {
	//beego.Router("/", &HomeController{})
	//beego.Run()

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@192.168.1.197:3306/py_zsd?charset=utf8", 30, 30) //注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	//ormer := orm.NewOrm()
	//raw := ormer.Raw("select * from user")
	//fmt.Println(raw)
}

func init() {

}