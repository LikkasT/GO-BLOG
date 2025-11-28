package middle_ware

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigInit(configname string, configpath string, configtype string) (*viper.Viper, error) {
	// 创建一个新的viper实例
	v := viper.New()
	v.SetConfigName(configname)
	v.SetConfigType(configtype)
	v.AddConfigPath(configpath)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	return v, nil
}
