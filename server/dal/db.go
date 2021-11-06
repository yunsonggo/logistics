package dal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"orderserve/model"
	"orderserve/model/backendModel"
)

var DB *xorm.Engine

func SqlEngine(conn string, driver string, show bool) {
	db, err := xorm.NewEngine(driver, conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n", err))
	}
	db.ShowSQL(show)
	err = db.Sync2(
		new(model.SmsCode),
		new(model.User),
		new(model.Goods),
		new(model.Cates),
		new(model.Give),
		new(model.Sale),
		new(model.Supplier),
		new(model.ShopCates),
		new(model.Evaluation),
		new(model.UserAddress),
		new(model.SaleGoods),
		new(model.Orders),
		new(backendModel.Manager),
		new(backendModel.BackendMenu),
		)
	if err != nil {
		panic(fmt.Sprintf("创建数据表出现异常:%v\n", err))
	}
	DB = db
	return
}
