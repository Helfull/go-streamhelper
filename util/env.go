package util

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func LoadDotEnv() (map[string]string, error) {

	var env map[string]string
	env, err := godotenv.Read()

	return env, err
}

func GetEnvStr(key string, fallback string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, ErrEnvVarEmpty
	}
	return v, nil
}

func GetEnvInt(key string, fallback int) (int, error) {
	s, err := GetEnvStr(key, "")
	if err != nil {
		return fallback, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback, err
	}
	return v, nil
}

func GetEnvBool(key string, fallback bool) (bool, error) {
	s, err := GetEnvStr(key, "false")
	if err != nil {
		return fallback, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return fallback, err
	}
	return v, nil
}
