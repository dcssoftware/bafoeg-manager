package configuration

import (
	"os"
	"strconv"
	"strings"

	safecast "github.com/ccoveille/go-safecast/v2"
)

func GetEnvOrDefaultString(name, defaultValue string) string {
	value := os.Getenv("APP_" + name)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvOrDefaultStringArray(name string, defaultValue []string) []string {
	value := os.Getenv("APP_" + name)
	if value == "" {
		return defaultValue
	}

	values := strings.Split(value, ",")

	var result []string = make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, strings.TrimSpace(v))
	}

	return result
}

func GetEnvOrDefaultInt(name string, defaultValue int) int {
	valueStr := os.Getenv("APP_" + name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func GetEnvOrDefaultUint(name string, defaultValue uint) uint {
	valueStr := os.Getenv("APP_" + name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		uintValue, uintValueErr := safecast.Convert[uint](value)
		if uintValueErr != nil {
			return 0
		}
		return uintValue
	}
	return defaultValue
}

func GetEnvOrDefaultBool(name string, defaultValue bool) bool {
	valueStr := os.Getenv("APP_" + name)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// func getEnvOrDefaultFloat(name string, defaultValue float64) float64 {
// 	valueStr := os.Getenv("APP_" + name)
// 	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
// 		return value
// 	}
// 	return defaultValue
// }
