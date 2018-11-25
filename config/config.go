package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 初始化配置文件方法
func (c *Config) initConfig() error {
	if c.Name != "" { // 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else { // 如果没有指定配置文件，则解析默认的盘【配置文件
		viper.AddConfigPath("config")
		viper.SetConfigName("snmp")
	}

	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")

	// viper解析配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}
