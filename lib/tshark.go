package lib

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	duration = 5
)

// Name returns name
func (t *TShark) Name(c *ReconConfig) string {
	return fmt.Sprintf("%s_tshark.pcap", c.InterfaceName)
}

// Run retrieves a sniffresult
func (t *TShark) Run(c *ReconConfig, out *bytes.Buffer) error {
	cmd := exec.Command(
		"tshark",
		"-I",
		"-i", c.InterfaceName,
		"-w", "-",
		"-a", fmt.Sprintf("duration:%d", duration),
	)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error running command: %+v", err)
	}
	return nil
}
