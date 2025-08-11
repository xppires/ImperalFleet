package config

import (
	"app/internal/common"
	"app/internal/interfaces"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// InitDB initMysql initializes the MySQL database connection
// It reads the database connection parameters from environment variables.
// It connects to the MySQL database and checks if the connection is successful.
func InitDB(configApp *ConfigApp) interfaces.DBStore {

	config := configApp.Database

	var dsn string

	switch config.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User, config.Password, config.Host, config.Port, config.DBName)
	case "local":
		return nil // Local repository does not require a database connection
	case "grpc":
		// gRPC does not require a database connection, return nil
		return nil
	// Add other database drivers as needed
	default:
		log.Fatal("unsupported database driver: %s", config.Driver)
	}

	fmt.Printf("Connecting to %s database at %s:%d...\n", config.Driver, config.Host, config.Port)
	fmt.Printf("DSN: %s\n", dsn)
	// Open a new database connectio

	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		log.Fatal("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(common.GetEnvIntOrDefault("DB_MAXOPEN", 25)) // Max open connections
	db.SetMaxIdleConns(common.GetEnvIntOrDefault("DB_MAXIDLE", 10)) // Max idle connections
	db.SetConnMaxLifetime(30 * time.Minute)                         // Max lifetime for a connection
	db.SetConnMaxIdleTime(10 * time.Minute)                         // Max idle time before closing

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging DB:", err)
	}
	var dbStore interfaces.DBStore
	dbStore = db
	return dbStore
}
