package utils

import (
	"os"
	"vm-init-utils/common"
)

func GetOSType() (string, error) {
	_, err := os.Stat(common.OSTypeFlagFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return common.Ubuntu, nil
		}
		return "", err
	}
	return common.Centos, nil
}
