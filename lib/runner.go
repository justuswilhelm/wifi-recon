package lib

import (
	"bytes"
	"fmt"
	"log"
)

// Run performs all reconnaisance operations
func (r *Recon) Run() error {
	r.Results = make([]ReconResult, len(r.Queue))
	for i, t := range r.Queue {
		outBuffer := new(bytes.Buffer)
		n := fmt.Sprintf("%T", t)
		log.Printf("Running Recon %s", n)
		r.Results[i] = ReconResult{
			Name: t.Name(r.Config),
		}
		if err := t.Run(r.Config, outBuffer); err != nil {
			r.Results[i].Error = err
			log.Printf("%s failed with: %s", n, err)
		} else {
			log.Printf("Resulting output size %d", outBuffer.Len())
			r.Results[i].Output = outBuffer
		}
		log.Printf("Finished running Recon %s", n)
	}
	return nil
}
