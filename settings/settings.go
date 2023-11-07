package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量 保存程序所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `yaml:"name"`
	Mode         string `yaml:"mode"`
	Version      string `yaml:"version"`
	Port         int    `yaml:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `mapstructure:"max_size"` // 下划线字段要使用mapstructure
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type MySQLConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	MaxConns     int    `mapstructure:"max_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config.yaml") // 指定配置文件
	//viper.AddConfigPath(".")           // 指定查找配置文件的路径

	// 配合远程配置中心使用
	//viper.SetConfigType("json")

	// 读取配置信息
	if err = viper.ReadInConfig(); err != nil {
		// 读取配置信息失败
		fmt.Printf("viper read config failed, err:%v\n", err)
		return err
	}

	// 将读取到的配置信息反序列化到 Conf 中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal config failed, err:%v\n", err)
		return err
	}
	fmt.Printf("AppConfig: %#v\n", Conf)
	fmt.Printf("LogConfig: %#v\n", Conf.LogConfig)
	fmt.Printf("MySQLConfig: %#v\n", Conf.MySQLConfig)
	fmt.Printf("RedisConfig: %#v\n", Conf.RedisConfig)

	viper.WatchConfig() // 监控配置文件变化

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		// 将读取到的配置信息反序列化到 Conf 中
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper unmarshal config failed, err:%v\n", err)
		}
	})
	return err
}
