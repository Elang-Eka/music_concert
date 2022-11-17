package configreader

import (
	"fmt"

	"github.com/spf13/viper"
)

type Interface interface {
	ReadConfig(cfg interface{})
	AllSettings() map[string]interface{}
}

type Options struct {
	Name string
	Type string
	Path string
}

type configReader struct {
	viper *viper.Viper
	opt   Options
}

func Init(opt Options) Interface {
	vp := viper.New()
	vp.SetConfigName(opt.Name)
	vp.SetConfigType(opt.Type)
	vp.AddConfigPath(opt.Path)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("Name :", opt.Name)
		fmt.Println("Type :", opt.Type)
		fmt.Println("Path :", opt.Path)
		panic(fmt.Errorf("fatal error found during reading file. err: %w", err))
	}

	c := &configReader{
		viper: vp,
		opt:   opt,
	}

	return c
}

func (c *configReader) ReadConfig(cfg interface{}) {
	if err := c.viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error found during unmarshaling config. err: %w", err))
	}
}

func (c *configReader) AllSettings() map[string]interface{} {
	return c.viper.AllSettings()
}
