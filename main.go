package main

import (
	"flag"
	"os"
	_ "os/exec"
    "io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
)

var (
	config string = ""
)

type ParachuteConfig struct {
	BackupHost         string `toml:"backup_host"`
	LocalTarballPrefix string `toml:"local_tarball_prefix"`
	OutputDirectory    string `toml:"output_directory"`
	Keep               int    `toml:"keep"`
	RemotePaths        string `toml:"remote_paths"`
	AdminEmail         string `toml:"admin_email"`
	PreBackup          string `toml:"pre_backup"`
}

func init() {
	flag.StringVar(&config, "config", "", "Path to config file (required)")
	flag.Parse()
	if config == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func LoadConfig() ParachuteConfig {
    var configobj ParachuteConfig
    data, err := ioutil.ReadFile(config)
    if err != nil {
        log.Errorf("Error reading config at %s: %s", config, err)
        os.Exit(1)
    }
    _, err = toml.Decode(string(data), &configobj)
    if err != nil {
        log.Errorf("Error parsing config at %s: %s", config, err)
        os.Exit(2)
    }
    return configobj
}

func main() {
	log.Infof("Starting parachute on config file %s", config)
    config := LoadConfig()
    log.Infof("config: %V", config)
}
