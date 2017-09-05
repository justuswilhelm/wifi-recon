package lib

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// RunWithTimeout runs a command and times out after t seconds
func RunWithTimeout(cmd *exec.Cmd, t int) error {
	done := make(chan error)
	go func() {
		err := cmd.Run()
		done <- err
	}()
	select {
	case err := <-done:
		return err
	case <-time.After(time.Second * time.Duration(t)):
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("Could not clean up process: %+v", err)
		}

		return fmt.Errorf("Timeout after %d seconds", t)
	}
}
