package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for migrate package
	"go.uber.org/zap"
)

type Config struct {
	Name            string
	Host            string
	Port            string
	User            string
	Password        string
	DisableSSL      bool
	MigrationFolder string
}

func Connect(config Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Host, config.Port, config.User, config.Password, config.Name)
	if config.DisableSSL {
		dsn += " sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		zap.L().Fatal(
			"Database connection failed",
			zap.Error(err),
			zap.String("user", config.User),
			zap.String("host", config.Host),
			zap.String("port", config.Port),
		)
	}
	if err := db.Ping(); err != nil {
		zap.L().Fatal("Can not ping database", zap.Error(err))
	}
	return db
}

func Close(db *sql.DB) {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		zap.L().Fatal("Database connection could not closed", zap.Error(err))
	}
}

func Migrate(db *sql.DB, config Config) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		zap.L().Error("could not start sql migration", zap.Error(err))
	}

	m, err := migrate.NewWithDatabaseInstance(config.MigrationFolder, config.Name, driver)
	if err != nil {
		zap.L().Error("migration failed", zap.Error(err))
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		zap.L().Error("Database synchronization failed", zap.Error(err))
	}
}
