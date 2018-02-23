package lib

import (
	"encoding/json"
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
	home := os.Getenv("HOME")
	if home == "" {
		log.Panicf("No $HOME in env")
	}
	var folderName string
	if r.Config.Name != "" {
		folderName = fmt.Sprintf("%s-%s", t, r.Config.Name)
	} else {
		folderName = t
	}
	return path.Join(home, "wifi-recon", "result", folderName)
}

func (r *Recon) ensurePath() error {
	path := r.getPath()
	log.Printf("Ensuring %s exists", path)
	return os.MkdirAll(path, os.ModePerm)
}

// WriteResults outputs recon results
func (r *Recon) WriteResults() error {
	outDir := r.getPath()

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
		path := path.Join(outDir, n)
		log.Printf("Writing result %s to %s", n, path)
		err := ioutil.WriteFile(path, result.Output.Bytes(), 0666)
		if err != nil {
			return fmt.Errorf("Error when writing to file: %+v", err)
		}
	}
	jsonFile := fmt.Sprintf("recon_result.json")
	jsonPath := path.Join(outDir, jsonFile)
	log.Printf("Writing results summary to %s", jsonPath)
	b, err := json.Marshal(r.Results)
	if err != nil {
		return fmt.Errorf("Error when marshalling results: %+v", err)
	}
	if err := ioutil.WriteFile(jsonPath, b, 0666); err != nil {
		return fmt.Errorf("Error when writing %s: %+v", jsonPath, err)
	}
	return nil
}
