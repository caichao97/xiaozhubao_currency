package Models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type UserModel struct {
	base baseModel
}
type User struct {
	Id int64
	Name string
}

func (user UserModel)init(){
	user.base.init()
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "User"
}