package Models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
)

type baseModel struct {
	dbType string
	dbClient string
	dbMaxIdle string
	dbMaxOpen string
}

func (base baseModel) init() {
	conf, err := config.NewConfig("ini", "app.conf")

	if err != nil {
		return
	}

	base.dbType = conf.String("defaultDb::dbType")

	base.dbClient = conf.String("defaultDb::dbClient")

	base.dbMaxIdle = conf.String("defaultDb::dbMaxIdle")

	base.dbMaxOpen = conf.String("defaultDb::dbMaxOpen")

	// set default database
	dbError :=orm.RegisterDataBase("default", base.dbType, base.dbClient)
	if dbError != nil {
		fmt.Print("db fail")
		return
	}

	maxIdle,_ :=strconv.Atoi(base.dbMaxIdle)

	orm.SetMaxIdleConns("default", maxIdle)

	maxOpen,_ :=strconv.Atoi(base.dbMaxOpen)

	orm.SetMaxOpenConns("default", maxOpen)
}