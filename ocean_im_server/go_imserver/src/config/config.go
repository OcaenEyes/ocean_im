package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	ServerIp      string `yaml:"serverIp"`
	ServerVersion string `yaml:"serverVersion"`
	Mysql         struct {
		DBAddress      string `yaml:"mysqlAddress"`
		DBPort         string `yaml:"mysqlPort"`
		DBUserName     string `yaml:"mysqlUsername"`
		DBPassWord     string `yaml:"mysqlPassword"`
		DBDatabaseName string `yaml:"mysqlDbname"`
		DBMaxOpenConns int    `yaml:"mysqlDbMaxOpenConns"`
		DBMaxIdleConns int    `yaml:"mysqlDbMaxIdleConns"`
		DBMaxLifeTime  int    `yaml:"mysqlDbMaxLifeTime"`
	}
	Redis struct {
		DBAddress     string `yaml:"redisAddress"`
		DBPort        string `yaml:"redisPort"`
		DBPassWord    string `yaml:"redisPass"`
		DBMaxIdle     int    `yaml:"redisMaxIdle"`
		DBMaxActive   int    `yaml:"redisMaxActive"`
		DBIdleTimeout int    `yaml:"redisIdleTimeout"`
	}
}

var Conf Config

func init() {
	bytes, err := ioutil.ReadFile("resources/conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &Conf)
	if err != nil {
		panic(err)
	}
}
