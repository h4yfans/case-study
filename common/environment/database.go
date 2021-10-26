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
		Name:            os.Getenv("POSTGRES_DB"),
		Host:            os.Getenv("POSTGRES_HOST"),
		Port:            os.Getenv("POSTGRES_PORT"),
		User:            os.Getenv("POSTGRES_USER"),
		Password:        os.Getenv("POSTGRES_PASSWORD"),
		DisableSSL:      getSSLMode(),
		MigrationFolder: getMigrationFolder(),
	}
}

func getSSLMode() bool {
	return strings.ToUpper(os.Getenv("POSTGRES_SSL_DISABLE")) == "TRUE"
}

func getMigrationFolder() string {
	if folder := os.Getenv("MIGRATION_FOLDER"); folder != "" {
		return folder
	}
	return DefaultMigrationFolder
}
