package main

import (
	"flag"
	"log"
	"os"

	"github.com/justuswilhelm/wifi-recon/lib"
)

var (
	// Interface to use
	Interface = ""
	// Name for dump files
	Name = ""
)

func init() {
	flag.StringVar(&Interface, "I", "", "Interface to use")
	flag.StringVar(&Name, "N", "", "Name to use for this dump")
}

func main() {
	flag.Parse()
	if Interface == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	directory, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	r := &lib.Recon{
		Config: &lib.ReconConfig{
			Name:          Name,
			InterfaceName: Interface,
			Path:          directory,
		},
		Queue: []lib.ReconTool{
			&lib.LocateMe{},
			&lib.TShark{},
			&lib.ArpScan{},
		},
	}
	if err := r.Run(); err != nil {
		log.Panicf("Error when runnning Recon: %+v", err)
	}
	if err := r.WriteResults(); err != nil {
		log.Panicf("Error when storing Results: %+v", err)
	}
}
