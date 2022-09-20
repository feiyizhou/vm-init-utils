package utils

import (
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
