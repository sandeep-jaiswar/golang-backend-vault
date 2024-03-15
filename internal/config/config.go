package config

import (
	"errors"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// LoadEnvFile loads environment variables from a YAML file.
func LoadEnvFile(path string) error {
	// Read the config file
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshal the config file
	var config map[string]string
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	// Set the environment variables
	for key, value := range config {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}

// ReadEnv reads an environment variable and returns its value as a string.
// It returns an error if the environment variable is not set.
func ReadEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("environment variable not set")
	}
	return value, nil
}

// ReadEnvInt reads an environment variable and returns its value as an integer.
// It returns an error if the environment variable is not set or if the value
// cannot be converted to an integer.
func ReadEnvInt(key string) (int, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return 0, errors.New("environment variable not set")
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ReadEnvBool reads an environment variable and returns its value as a boolean.
// It returns an error if the environment variable is not set or if the value
// is not a valid boolean.
func ReadEnvBool(key string) (bool, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return false, errors.New("environment variable not set")
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, err
	}
	return value, nil
}
