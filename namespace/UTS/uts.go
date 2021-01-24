package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
 * UTS Namespace主要用来隔离node name和domain name。在每个UTS Namespace中，每个Namespace都允许有自己的hostname
 */

func main() {
	// 指定被fork出来的新进程的初始命令
	cmd := exec.Command("sh")

	// 创建uts namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
