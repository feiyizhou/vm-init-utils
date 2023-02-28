package utils

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// ExecCMDWithResult 执行命令并返回结果
func ExecCMDWithResult(name string, args []string, duration int) ([]byte, error) {
	// 5秒超时
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	command := exec.CommandContext(ctx, name, args...)
	closer, err := command.StdoutPipe()
	defer func() {
		cancelFunc()
		_ = closer.Close()
		_ = command.Wait()
	}()
	if err != nil {
		return nil, err
	}
	err = command.Start()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(closer)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

// ExecCmd 执行命令
func ExecCmd(cmdName string, args []string) error {
	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
