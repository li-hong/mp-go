package main

import (
	_ "mp-go/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "zan:hIszan2016@tcp(liangyibang.com:3001)/zanwh_mp?charset=utf8")
}
