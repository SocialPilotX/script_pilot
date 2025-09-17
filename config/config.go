package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	App struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"app"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	GeminiKey string `yaml:"gemini_key"`
}

var Config AppConfig

func init() {
	_ = godotenv.Load(".env")
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	fileName := fmt.Sprintf("config/config.%s.yaml", env)
	raw, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	parsed := os.ExpandEnv(string(raw))
	if err := yaml.NewDecoder(strings.NewReader(parsed)).Decode(&Config); err != nil {
		log.Fatalf("failed to parse config YAML: %v", err)
	}
}
