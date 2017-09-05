package lib

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Run Arp Scan
func (l *LocateMe) Run(c *ReconConfig, out *bytes.Buffer) error {
	cmd := exec.Command(
		"locateme",
		"-l",
	)
	cmd.Stdout = out
	err := RunWithTimeout(cmd, 4)
	if err != nil {
		return fmt.Errorf("Error when running locateme: %+v", err)
	}
	return nil
}

// Name of locateme
func (l *LocateMe) Name(c *ReconConfig) string {
	return "locateme"
}
