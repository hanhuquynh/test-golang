package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id      int64  `xorm:"pk autoincr"`
	Name    string `xorm:"varchar(255)"`
	Age     int
	Created int64
	Updated int64
}

func ConnectDb() (*xorm.Engine, error) {
	mySQLUserName := os.Getenv("MYSQL_USERNAME")
	mySQLPassword := os.Getenv("MYSQL_PASSWORD")
	mySQLHost := os.Getenv("MYSQL_HOST")
	mySQLPort := os.Getenv("MYSQL_PORT")
	mySQLDbName := os.Getenv("MYSQL_DBNAME")
	// "username:password@tcp(127.0.0.1:3306)/test"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mySQLUserName, mySQLPassword, mySQLHost, mySQLPort, mySQLDbName)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	if err := engine.Ping(); err != nil {
		return nil, err
	}
	return engine, nil
}

func CreateTableUser() error {
	b, err := d.IsTableExist(&User{})
	if err != nil {
		return err
	}
	if b {
		if err := d.Sync2(&User{}); err != nil {
			return err
		}
	} else {
		if err := d.CreateTables(&User{}); err != nil {
			return err
		}
	}
	return nil
}
