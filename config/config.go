package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// InitConfig inits config
func InitConfig() (Manager, func(), error) {
	//ServiceName = "RAMME-TEMPLATE"
	confManager := NewMutexConfigManager(LoadConfig(File))

	// Watch the file for modification and update the config manager with the new config when it's available
	watcher, err := WatchFile(File, time.Second, func() {
		conf := LoadConfig(File)
		confManager.Set(conf)
	})
	if err != nil {
		return nil, nil, err
	}
	return confManager, func() { watcher.Close() }, nil
}

// LoadConfig loads config from file loader
func LoadConfig(configFile string) *Config {
	conf := &Config{}
	configData, err := ioutil.ReadFile(filepath.Clean(configFile))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(configData, conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return conf
}
