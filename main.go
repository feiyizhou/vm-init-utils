package main

import (
	"log"
	"vm-init-utils/cmd"
)

func main() {
	err := cmd.NewChangeCmd().Execute()
	if err != nil {
		log.Fatalf("Os init failed, err : %v \n", err)
	}
}
