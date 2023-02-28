package utils

import (
	"fmt"
	"strings"
)

var (
	WinGBKCode = "936"
	WinGBKName = "gbk"
	WinUTFCode = "65001"
	WinUTFName = "utf"
)

func GetWinDecodeType() (string, error) {
	result, err := ExecCMDWithResult("chcp", nil, 10)
	if err != nil {
		return "", err
	}
	switch true {
	case strings.Contains(string(result), WinGBKCode):
		return WinGBKName, nil
	case strings.Contains(string(result), WinUTFCode):
		return WinUTFName, nil
	default:
		return "", fmt.Errorf("Unknown decode type: %s ", string(result))
	}
}
