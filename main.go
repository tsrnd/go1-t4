package main

import (
	"fmt"
	_ "github.com/goweb4/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	driverName 	:= beego.AppConfig.String("dbconnection")
	dbUser			:= beego.AppConfig.String("dbusername")
	dbPassword	:= beego.AppConfig.String("dbpassword")
	dbDatabase	:= beego.AppConfig.String("dbdatabase")
	dbHost			:= beego.AppConfig.String("dbhost")
	dbInfor			:= fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbDatabase,
	)
	maxIdle 		:= 30
	maxConn 		:= 30

	orm.RegisterDriver(driverName, orm.DRPostgres)
	orm.RegisterDataBase("default", driverName, dbInfor, maxIdle, maxConn)
}

func main() {
	beego.SetStaticPath("/static","static")
	o := orm.NewOrm()
	orm.Debug = true
	o.Using("default") // Using default, can use other database

	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}
