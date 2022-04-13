package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"new-project/pkg/util"
	"time"
)

type Config struct {
	Service *Service  `yaml:"Service"`
	Logger  *Logger   `yaml:"Logger"`
	DB      *Database `yaml:"DB"`
	Redis   *Redis    `yaml:"Redis"`
	Oss     *Oss      `yaml:"Oss"`
	Jwt     *Jwt      `yaml:"Jwt"`
}

var config Config

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

func GetJwt() *Jwt {
	return config.Jwt
}

func GetJwtExpire() time.Duration {
	return time.Duration(config.Jwt.Expire) * time.Hour
}

func GetJwtSecret() []byte {
	return []byte(config.Jwt.Secret)
}
