package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("ls", "1000000")
	//cmd := exec.Command("sleep", "1")
	cmd.Start()
	done := make(chan struct{})
	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Printf("wait done return err:%v", err)
		}
		status := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitStatus := status.ExitStatus()
		signaled := status.Signaled()
		signal := status.Signal()
		fmt.Println("Error:", err)
		if signaled {
			fmt.Println("Signal:", signal)
		} else {
			fmt.Println("Status:", exitStatus)
		}
		close(done)
	}()
	//cmd.Process.Kill()
	<-done
	fmt.Printf("process another thing\n")
}
