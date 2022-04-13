package config

// Service 应用服务配置
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

// Logger 日志服务配置
type Logger struct {
	DebugFile  string `yaml:"DebugFile"`
	AccessFile string `yaml:"AccessFile"`
	MaxSize    int    `yaml:"MaxSize"`
	MaxAge     int    `yaml:"MaxAge"`
}

// Database 数据库服务配置
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

// Redis redis服务配置
type Redis struct {
	Host      string `yaml:"Host"`
	Password  string `yaml:"Password"`
	DefaultDB int    `yaml:"DefaultDB"`
}

// Oss 阿里云OSS服务配置
type Oss struct {
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	BucketName      string `yaml:"BucketName"`
}

// Jwt 用户 jwt token 配置
type Jwt struct {
	Expire int    `yaml:"Expire"`
	Secret string `yaml:"Secret"`
}
