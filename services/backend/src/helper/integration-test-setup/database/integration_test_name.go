package database

import "fmt"

func CreateintegrationDatabaseName(hashID uint64) string {
	return fmt.Sprintf("integration_test_%d", hashID)
}
