package helper

import (
	"os"
	"strconv"
	"strings"
)

// getEnv return env variable or default value provided
func GetEnvStr(name, defaultV string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return defaultV
}

func GetEnvBool(name string, defaultV bool) bool {
	if value, ok := os.LookupEnv(name); ok {
		return strings.ToLower(value) == "true"
	}

	return defaultV
}

func GetEnvInt64(name string, defaultV int64) int64 {
	if value, ok := os.LookupEnv(name); ok {
		if v, err := strconv.ParseInt(value, 10, 64); err == nil {
			return v
		}
	}

	return defaultV
}
