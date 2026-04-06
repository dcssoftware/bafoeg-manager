package numbers

import "strconv"

func ConvertStringToNilInt(value string) *int {
	intVal, intValErr := strconv.Atoi(value)

	if intValErr != nil {
		return nil
	}

	return &intVal
}

func ConvertStringToInt(value string, defaultValue int) int {
	intVal, intValErr := strconv.Atoi(value)

	if intValErr != nil {
		return defaultValue
	}

	return intVal
}
