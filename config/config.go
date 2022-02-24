package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Service struct {
	DebugMode   bool   `yaml:"debug_mode"`
	MaxPageSize int    `yaml:"max_page_size"`
	MinPageSize int    `yaml:"min_page_size"`
	Port        string `yaml:"port"`
}

type Logger struct {
	DebugFile  string `yaml:"debug_file"`
	AccessFile string `yaml:"access_file"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
}
type Config struct {
	Service Service `yaml:"service"`
	Logger  Logger  `yaml:"logger"`
}

var config Config

func GetService() Service {
	return config.Service
}

func GetLogger() Logger {
	return config.Logger
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
