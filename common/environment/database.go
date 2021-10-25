package environment

import (
	"os"
	"strings"

	"github.com/h4yfans/case-study/common/db"
)

var (
	DefaultMigrationFolder = "file://db/migrations"
)

func Database() db.Config {
	return db.Config{
		Name:            os.Getenv("DB_NAME"),
		Host:            os.Getenv("DB_HOST"),
		Port:            os.Getenv("DB_PORT"),
		User:            os.Getenv("DB_USERNAME"),
		Password:        os.Getenv("DB_PASSWORD"),
		DisableSSL:      getSSLMode(),
		MigrationFolder: getMigrationFolder(),
	}
}

func getSSLMode() bool {
	return strings.ToUpper(os.Getenv("DB_SSL_DISABLE")) == "TRUE"
}

func getMigrationFolder() string {
	if folder := os.Getenv("MIGRATION_FOLDER"); folder != "" {
		return folder
	}
	return DefaultMigrationFolder
}
