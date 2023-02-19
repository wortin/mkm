package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf *viper.Viper

func init() {
	conf = viper.New()
	conf.SetConfigFile("conf.yaml")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

var DB *gorm.DB

func init() {
	var err error
	user := conf.GetString("mysql.user")
	pd := conf.GetString("mysql.password")
	host := conf.GetString("mysql.host")
	port := conf.GetString("mysql.port")
	db := conf.GetString("mysql.database")
	args := conf.GetString("mysql.args")

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pd, host, port, db, args)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 建表
	err = DB.AutoMigrate(&HanZi{})
	if err != nil {
		panic(err)
	}

}
