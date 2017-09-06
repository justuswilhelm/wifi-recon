package lib

import (
	"bytes"
)

// ReconTool defines the shared operations a tool has to offer
type ReconTool interface {
	Name(config *ReconConfig) string
	Run(config *ReconConfig, out *bytes.Buffer) error
}

// TShark stores airport sniff results
type TShark struct {
}

// ArpScan perfomrs arpscan
type ArpScan struct {
}

// LocateMe outputs geolocation
type LocateMe struct {
}

// Recon stores all types of results
type Recon struct {
	Queue   []ReconTool
	Results []ReconResult
	Config  *ReconConfig
}

// ReconConfig stores common parameters such as output path
type ReconConfig struct {
	InterfaceName string
	Name          string
	Path          string
}

// ReconResult stores the output of a recon tool's result
type ReconResult struct {
	Output *bytes.Buffer
	Error  error
	Name   string
}
