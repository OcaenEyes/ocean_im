package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"oceanim/go_server/src/config"
	"time"
)

var Db *sql.DB

func initMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Conf.Mysql.DBUserName,
		config.Conf.Mysql.DBPassWord,
		config.Conf.Mysql.DBAddress,
		config.Conf.Mysql.DBPort,
		"mysql")
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		//fmt.Println("open mysql failed", err)
		//return
		panic(err)
	}
	createDbSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s ;", config.Conf.Mysql.DBDatabaseName)
	_, err = database.Exec(createDbSql)
	if err != nil {
		panic(err)
	}
	database.Close()
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Conf.Mysql.DBUserName,
		config.Conf.Mysql.DBPassWord,
		config.Conf.Mysql.DBAddress,
		config.Conf.Mysql.DBPort,
		config.Conf.Mysql.DBDatabaseName)

	database, err = sql.Open("mysql", dsn)
	Db = database
	Db.SetMaxOpenConns(config.Conf.Mysql.DBMaxOpenConns)
	Db.SetMaxIdleConns(config.Conf.Mysql.DBMaxIdleConns)
	Db.SetConnMaxLifetime(time.Duration(config.Conf.Mysql.DBMaxLifeTime) * time.Second)
}
