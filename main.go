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
)

func init() {
	flag.StringVar(&Interface, "I", "", "Interface to use")
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
