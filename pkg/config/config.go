package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"new-project/pkg/util"
)

type Service struct {
	DebugMode            bool     `yaml:"DebugMode"`
	MaxPageSize          int      `yaml:"MaxPageSize"`
	MinPageSize          int      `yaml:"MinPageSize"`
	UploadChunkSize      int      `yaml:"UploadChunkSize"`
	UploadImgMaxSize     float64  `yaml:"UploadMaxImgSize"`
	UploadVideoMaxSize   float64  `yaml:"UploadVideoMaxSize"`
	Host                 string   `yaml:"Host"`
	AppName              string   `yaml:"AppName"`
	UploadSavePath       string   `yaml:"UploadSavePath"`
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

type Oss struct {
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	BucketName      string `yaml:"BucketName"`
}

type Config struct {
	Service *Service  `yaml:"Service"`
	Logger  *Logger   `yaml:"Logger"`
	DB      *Database `yaml:"DB"`
	Redis   *Redis    `yaml:"Redis"`
	Oss     *Oss      `yaml:"Oss"`
}

var config Config

func GetService() *Service {
	return config.Service
}

func (service *Service) GetChunkSize() uint {
	return uint(service.UploadChunkSize * util.MB)
}

func GetLogger() *Logger {
	return config.Logger
}

func GetDB() *Database {
	return config.DB
}

func GetRedis() *Redis {
	return config.Redis
}

func GetOss() *Oss {
	return config.Oss
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
