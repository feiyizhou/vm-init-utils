package utils

import (
	"fmt"
	"strings"
)

const (
	GBKDecodeCode = "936"
	UTFDecodeCode = "65001"
)

func GetWinDecodeType() (string, error) {
	result, err := ExecCMDWithResult("chcp", nil, 10)
	if err != nil {
		return "", err
	}
	switch true {
	case strings.Contains(string(result), GBKDecodeCode):
		return GBKDecodeCode, nil
	case strings.Contains(string(result), UTFDecodeCode):
		return UTFDecodeCode, nil
	default:
		return "", fmt.Errorf("Unknown decode type: %s ", string(result))
	}
}
