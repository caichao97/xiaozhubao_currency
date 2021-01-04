package Models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Workids struct {
	Id int64
	WorkId int64
}

func (WorkIds *Workids) TableName() string {
	return "Workids"
}

func init() {
	orm.RegisterModel(new(Workids))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/currency_service?charset=utf8")
}