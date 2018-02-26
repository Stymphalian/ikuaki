package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	path, err := exec.LookPath("../agent/agent")
	if err != nil {
		log.Fatal(err)
	}

	defaultPgid := 0
	for i := 0; i < 5; i++ {
		cmd := exec.Command(path, fmt.Sprintf("--name=name%d", i))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
			Pgid:    defaultPgid,
		}
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
			defaultPgid, err = syscall.Getpgid(cmd.Process.Pid)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	time.Sleep(20 * time.Second)
}
