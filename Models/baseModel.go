package Models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type baseModel struct {
	dbType string
	dbClient string
}

func (base baseModel)init() {
	conf, err := config.NewConfig("ini", "app.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	base.dbType = conf.String("defaultDb::dbType")

	base.dbClient = conf.String("defaultDb::dbClient")

	// set default database
	_ = orm.RegisterDataBase("default", base.dbType, base.dbClient, 30)
	// create table
	_ = orm.RunSyncdb("default", false, true)
}