package config

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	BaseDir string `json:"BaseDir"`
}

func Load(filePath string) *Config {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	cf, err := os.Open(filepath.Join(usr.HomeDir, filePath))
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
		return nil
	} else if os.IsNotExist(err) {
		return &Config{filepath.Join(usr.HomeDir, ".todo")}
	}
	defer cf.Close()

	jsonParser := json.NewDecoder(cf)
	var config Config
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
