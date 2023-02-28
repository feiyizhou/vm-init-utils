package utils

import (
	"fmt"
	"log"
)

func DoOrDieWithMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, err: %v", msg, err)
	}
}

func DieWithMsg(flag bool, msg string) {
	if flag {
		log.Fatalln(msg)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func MadeErr(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s, err: %v", msg, err)
	} else {
		return fmt.Errorf("%s", msg)
	}
}
