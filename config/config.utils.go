package config

import (
	"os"
	"strconv"
)

func getVarEnvAsStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getVarEnvAsInt(key string, defaultVal int) int {
	valueStr := getVarEnvAsStr(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
