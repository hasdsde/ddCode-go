package global

import (
	logger "ddCode-server/pkg/zlogs"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func parse(configFile string) (*Config, error) {
	c := new(Config)
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Parse(configFile, configModeFile *string) (*Config, error) {
	filename, _ := filepath.Abs(*configFile)
	c, err := parse(filename)
	if err != nil {
		fmt.Printf("解析配置文件失败:%v\n", err)
		return c, err
	}
	// 读取子配置文件
	li := strings.LastIndex(*configFile, ".")
	if *configModeFile == "" {
		*configModeFile = c.Mode
	}
	filename, _ = filepath.Abs((*configFile)[:li] + "-" + *configModeFile + (*configFile)[li:])
	cm, err := parse(filename)
	if err != nil {
		fmt.Printf("解析子配置文件失败:%v\n", err)
		return c, err
	}
	err = c.Together(cm)
	if err != nil {
		fmt.Printf("合并配置文件失败:%v\n", err)
		return c, err
	}
	return c, nil
}

type Config struct {
	Mode   string           `yaml:"Mode"`
	Host   string           `yaml:"Host"`
	Port   int              `yaml:"Port"`
	Mysql  MysqlInfo        `yaml:"Mysql"`
	Log    logger.LogConfig `yaml:"Log"`
	Jwt    JwtConf          `yaml:"Jwt"`
	Casbin CasbinConf       `yaml:"Casbin"`
	Redis  RedisConf        `yaml:"Redis"`
	Oss    OssConf          `yaml:"Oss"`
}
type CasbinConf struct {
	FilePath string `yaml:"FilePath"`
}
type MysqlInfo struct {
	Host     string `yaml:"Host"`
	Database string `yaml:"Database"`
	User     string `yaml:"User"`
	Pass     string `yaml:"Pass"`
	Port     string `yaml:"Port"`
}

type RedisConf struct {
	Addr string `yaml:"Addr"`
	Pass string `yaml:"Pass"`
	DB   int    `yaml:"DB"`
}

type OssConf struct {
	AccessKey string `yaml:"AccessKey"`
	SecretKey string `yaml:"Secret"`
	Bucket    string `yaml:"Bucket"`
	Url       string `yaml:"Url"`
}

type JwtConf struct {
	Signature string        `yaml:"Signature"`
	Duration  time.Duration `yaml:"Duration"`
}

func (c *Config) Together(sub *Config) (err error) {

	// 嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达	嗯，马达马达
	var copyCopy func(src, dst reflect.Value)
	copyCopy = func(src, dst reflect.Value) {

		t := src.Type()
		v := src

		sv := dst

		for i := 0; i < t.NumField(); i++ {
			tt := t.Field(i)

			vv := v.Field(i)
			svv := sv.Field(i)

			if tt.Tag.Get("config") == "ignore" {
				continue
			}

			setValue := func(v, sv reflect.Value) {
				if v.CanSet() {
					v.Set(sv)
				} else {
					panic("不能设置值")
				}
			}

			switch tt.Type.Kind() {
			case reflect.String:
				if svv.String() != "" && vv.String() == "" {
					fmt.Println("替换")
					setValue(vv, svv)
				}
			case reflect.Int:
				if vv.Int() == 0 && svv.Int() != 0 {
					fmt.Println("替换")
					setValue(vv, svv)
				}
			case reflect.Int64: // other
				if vv.Int() == 0 && svv.Int() != 0 {
					fmt.Println("替换")
					setValue(vv, svv)
				}
			case reflect.Bool:
				if !vv.Bool() && svv.Bool() {
					fmt.Println("替换")
					setValue(vv, svv)
				}
			case reflect.Pointer:
				if svv.IsNil() {
					continue
				}
				vv = v.Elem()
				svv = svv.Elem()
				fallthrough
			case reflect.Struct:
				copyCopy(vv, svv)
			default:
				fmt.Println(tt.Name, tt.Type.Kind(), tt.Tag.Get("config"), tt.Tag.Get("yaml"))
				panic("未添加的类型，需要修改代码")
			}
		}

	}

	copyCopy(reflect.ValueOf(c).Elem(), reflect.ValueOf(sub).Elem())

	return nil
}
