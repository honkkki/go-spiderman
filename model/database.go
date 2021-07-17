package model

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

// Model base model
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func InitDB(iniPath string) {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	cfg, err := ini.Load(iniPath)
	if err != nil {
		log.Fatalf("parse ini file fail: %v", err)
	}

	sec, _ := cfg.GetSection("database")

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(200)
}
