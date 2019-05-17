package main

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	 _ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string
	Profile *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id int
	Age int16
	User *User `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
	orm.RegisterDataBase("default", "mysql", "root:123456@/learn-beego?charset=utf8")
}

func main() {
	o := orm.NewOrm()

	fmt.Println("HI!!!")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM user WHERE name = ?", "Phan Dinh Man").Values(&maps)
	fmt.Println(len(maps))
	if err == nil && num > 0 {
		fmt.Println(maps[0]["ID"])
	}
}
