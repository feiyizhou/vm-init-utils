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
	Hostname string         `json:"hostname" mapstructure:"hostname"`
	PASSWD   string         `json:"passwd" mapstructure:"passwd"`
	Network  *NetworkConfig `json:"network" mapstructure:"network"`
}

type NetworkConfig struct {
	Name    string `json:"name" mapstructure:"name"`
	MACAddr string `json:"macAddr" mapstructure:"macAddr"`
	IPAddr  string `json:"ipAddr" mapstructure:"ipAddr"`
	NETMASK string `json:"netmask" mapstructure:"netmask"`
	GATEWAY string `json:"gateway" mapstructure:"gateway"`
	DNS1    string `json:"dns1" mapstructure:"dns1"`
	DNS2    string `json:"dns2" mapstructure:"dns2"`
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
