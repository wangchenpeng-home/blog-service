package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp * viper.Viper
}

// NewSetting 初始化本项目配置的基本属性
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}
