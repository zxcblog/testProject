package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Service struct {
	DebugMode            bool     `yaml:"DebugMode"`
	MaxPageSize          int      `yaml:"MaxPageSize"`
	MinPageSize          int      `yaml:"MinPageSize"`
	Host                 string   `yaml:"Host"`
	UploadSavePath       string   `yaml:"UploadSavePath"`
	UploadImgMaxSize     float64  `yaml:"UploadMaxImgSize"`
	UploadVideoMaxSize   float64  `yaml:"UploadVideoMaxSize"`
	UploadImageAllowExts []string `yaml:"UploadImageAllowExts"`
	UploadVideoAllowExts []string `yaml:"UploadVideoAllowExts"`
}

type Logger struct {
	DebugFile  string `yaml:"DebugFile"`
	AccessFile string `yaml:"AccessFile"`
	MaxSize    int    `yaml:"MaxSize"`
	MaxAge     int    `yaml:"MaxAge"`
}

type Database struct {
	DBType       string `yaml:"DBType"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	Port         string `yaml:"Port"`
	DBName       string `yaml:"DBName"`
	TablePrefix  string `yaml:"TablePrefix"`
	Charset      string `yaml:"Charset"`
	ParseTime    bool   `yaml:"ParseTime"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}

type Redis struct {
	Host      string `yaml:"Host"`
	Password  string `yaml:"Password"`
	DefaultDB int    `yaml:"DefaultDB"`
}

type Config struct {
	Service Service  `yaml:"Service"`
	Logger  Logger   `yaml:"Logger"`
	DB      Database `yaml:"DB"`
	Redis   Redis    `yaml:"Redis"`
}

var config Config

func GetService() Service {
	return config.Service
}

func GetLogger() Logger {
	return config.Logger
}

func GetDB() Database {
	return config.DB
}

func GetRedis() Redis {
	return config.Redis
}

func InitConfig(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("配置文件读取失败:%+v", err.Error()))
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败:%+v", err.Error()))
	}
}
