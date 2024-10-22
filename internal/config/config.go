package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	configPathKey = "CONFIG_PATH"

	envKey                = "ENV"
	envDevValue           = "dev"
	envProdValue          = "prod"
	defaultEnv            = envProdValue
	defaultDevConfigPath  = "./config-dev.yaml"
	defaultProdConfigPath = "./config-prod.yaml"
)

type Config struct {
	RuntimeDirPath          string        `yaml:"runtime_dir_path"`
	BorzoiConfigPath        string        `yaml:"borzoi_config_path"`
	ServerHost              string        `yaml:"server_host"`
	ServerPort              string        `yaml:"server_port"`
	RedisHost               string        `yaml:"redis_host"`
	RedisPort               string        `yaml:"redis_port"`
	RedisPassword           string        `yaml:"redis_password"`
	RedisDB                 int           `yaml:"redis_db"`
	RequestTimeout          time.Duration `yaml:"request_timeout"`
	EnableWatcher           bool          `yaml:"enable_watcher"`
	MaxPortSearchIterations int           `yaml:"max_port_search_iterations"`
	DefaultPortRangeMin     int           `yaml:"default_port_range_min"`
	DefaultPortRangeMax     int           `yaml:"default_port_range_max"`
	NginxSitesAvailablePath string        `yaml:"nginx_sites_available_path"`
	NginxSitesEnabledPath   string        `yaml:"nginx_sites_enabled_path"`
	HostsFilePath           string        `yaml:"hosts_file_path"`
}

func findConfig() (string, error) {
	configPath, foundInEnv := os.LookupEnv(configPathKey)
	if foundInEnv {
		return configPath, nil
	}

	env, foundInEnv := os.LookupEnv(envKey)
	if !foundInEnv {
		env = defaultEnv
	}

	switch env {
	case envDevValue:
		return defaultDevConfigPath, nil
	case envProdValue:
		return defaultProdConfigPath, nil
	}

	return "", fmt.Errorf(
		"can't find config, try setting %s env var or make sure there are configs on default paths: %s and %s",
		configPathKey, defaultDevConfigPath, defaultProdConfigPath,
	)
}

func NewConfig() Config {
	configPath, err := findConfig()
	if err != nil {
		log.Fatalf("failed to find config: %v", err)
	}

	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %v", err)
	}

	return config
}
