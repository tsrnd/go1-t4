package main

import (
	"fmt"
	_ "github.com/goweb4/routers"
	"github.com/astaxie/beego"
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
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // Using default, can use other database

	beego.Run()
}
