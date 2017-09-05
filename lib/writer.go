package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func (r *Recon) getPath() string {
	// RFC3339
	t := time.Now().Format(time.RFC3339)
	return path.Join("result", t)
}

func (r *Recon) ensurePath() error {
	path := r.getPath()
	log.Printf("Ensuring %s exists", path)
	return os.MkdirAll(path, os.ModePerm)
}

// WriteResults outputs recon results
func (r *Recon) WriteResults() error {
	if len(r.Results) == 0 {
		return fmt.Errorf("There are no results to write")
	}
	if err := r.ensurePath(); err != nil {
		return fmt.Errorf("Error when ensuring path: %+v", err)
	}
	for _, result := range r.Results {
		if result.Error != nil {
			log.Printf("Skipping %s as it yielded an error", result.Name)
			continue
		}
		n := result.Name
		out := r.getPath()
		path := path.Join(out, n)
		log.Printf("Writing result %s to %s", n, path)
		err := ioutil.WriteFile(path, result.Output.Bytes(), 0666)
		if err != nil {
			return fmt.Errorf("Error when writing to file: %+v", err)
		}
	}
	return nil
}
