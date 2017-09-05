package lib

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Run Arp Scan
func (a *ArpScan) Run(c *ReconConfig, out *bytes.Buffer) error {
	cmd := exec.Command(
		"arp-scan",
		"-l",
	)
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error when running arp-scan: %+v", err)
	}
	return nil
}

// Name of scanner
func (a *ArpScan) Name(c *ReconConfig) string {
	return fmt.Sprintf("%s_arpscan", c.InterfaceName)
}
