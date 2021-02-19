package main

import (
	"project/models"
	_ "project/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:kredit@66@/FirstBee?charset=utf8")

	//orm.RegisterDataBase("default", "mysql", "database/orm_test.db")

	orm.RegisterModel(new(models.Article))

}
func main() {
	beego.Run()
}
