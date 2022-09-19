package utils

import (
	"log"
	"os"
	"os/exec"
)

func ExecShell(args []string, in, out, erf *os.File) error {
	command := exec.Command("sh", args...)
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	if erf == nil {
		erf = os.Stderr
	}
	command.Stdin = in
	command.Stdout = out
	command.Stderr = erf
	return command.Run()
}

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
