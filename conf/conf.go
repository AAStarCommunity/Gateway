package conf

import (
	"sync"
)

var once sync.Once

type Conf struct {
	Node struct {
		Host string
	}
	Service struct {
		Port int
	}
}

var conf *Conf

// Get 读取配置
// 默认从配置文件取，如果配置文件中的db节点内容为空，则从环境变量取
// 如果配置文件不存在，则db从环境变量取，其他值使用默认值
func Get() *Conf {
	once.Do(func() {
		if conf == nil {
			filePath := getConfFilePath()
			conf = getConfiguration(filePath)
		}
	})
	return conf
}
