package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config 定义配置信息结构体，从配置文件读入
type Config struct {
	Host   string `json:"host"` // 监听的http地址
	Port   string `json:"port"` // 监听的http端口
	WebDir string `json:"web"`  // web静态文件所在的目录
}

// LoadConfig 读入配置文件
func LoadConfig(file string) (c *Config, err error) {
	// 将文件读到内存中，为一个切片类型
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// json解析切片数据，反序列化到结构体中
	c = &Config{}
	err = json.Unmarshal(data, c)
	fmt.Println(*c, err)
	return c, err
}
