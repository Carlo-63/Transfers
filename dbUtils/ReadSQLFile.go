package dbUtils

import (
	"os"
)

func ReadSQLFile(filePath string) (string, error) {
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(sqlBytes), nil
}
