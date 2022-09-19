package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"vm-init-utils/common"
	"vm-init-utils/utils"
)

type SysConfig struct {
	Hostname string        `json:"hostname" mapstructure:"hostname"`
	Network  networkConfig `json:"network" mapstructure:"network"`
}

type networkConfig struct {
	Name    string `json:"name" mapstructure:"name"`
	Addr    string `json:"addr" mapstructure:"addr"`
	Gateway string `json:"gateway" mapstructure:"gateway"`
	Mask    string `json:"mask" mapstructure:"mask"`
	DNS     string `json:"dns" mapstructure:"dns"`
	MAC     string `json:"mac" mapstructure:"mac"`
}

func GetSystemConf() *SysConfig {
	conf, err := initSystemConf()
	utils.CheckErr(err)
	utils.DieWithMsg(conf == nil, "System configuration is not exist")
	return conf
}

func initSystemConf() (*SysConfig, error) {
	viper.AddConfigPath(common.YamlConfigHomePath)
	viper.SetConfigName(common.YamlConfigName)
	viper.SetConfigType(common.YamlConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	sysConfMap := viper.GetStringMap(common.YamlSysConfigKey)
	if sysConfMap == nil {
		return nil, fmt.Errorf("Etcd config is not exist ")
	}
	sysConfig := new(SysConfig)
	err = ParseInterface2Struct(sysConfMap, &sysConfig)
	if err != nil {
		return nil, fmt.Errorf("Parse confStr : %+v to struct err , err : %+v ", sysConfMap, err)
	}
	return sysConfig, err
}

// ParseInterface2Struct ...
func ParseInterface2Struct(in interface{}, out interface{}) error {
	kind := reflect.TypeOf(in).Kind()
	if reflect.Map == kind {
		err := mapstructure.Decode(in, out)
		if err != nil {
			return err
		}
	} else if reflect.String == kind {
		err := json.Unmarshal([]byte(in.(string)), out)
		if err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("Can not parse this type : %s ! ", kind.String()))
	}
	return nil
}
