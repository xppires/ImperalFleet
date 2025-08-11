package config

import (
	"app/internal/common"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  string
	GrpcAddr string // gRPC address for remote repository
}
type ServerConfig struct {
	Addr     string
	GrpcAddr string
}
type JWTConfig struct {
	Key string
}

type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

func LoadConfig() (*AppConfig, error) {

	databaseConfig := DatabaseConfig{
		Driver:   common.GetEnvOrDefault("DB_DRIVER", "mysql"),
		Host:     common.GetEnvOrDefault("DB_HOST", "localhost"),
		Port:     common.GetEnvIntOrDefault("DB_PORT", 3306),
		DBName:   common.GetEnvOrDefault("DB_DATABASE", "task_api"),
		User:     common.GetEnvOrDefault("DB_USERNAME", "root"),
		Password: common.GetEnvOrDefault("DB_PASSWORD", ""),
		SSLMode:  common.GetEnvOrDefault("DB_SSL_MODE", "disable"),
		GrpcAddr: common.GetEnvOrDefault("DB_GRPC_ADDR", ":7070"),
	}
	serverConfig := ServerConfig{
		Addr:     common.GetEnvOrDefault("APP_URL", ":8080"),
		GrpcAddr: common.GetEnvOrDefault("DB_GRPC_ADDR", ":9090"),
	}
	jWTConfig := JWTConfig{Key: common.GetEnvOrDefault("APP_SECRET", "disable")}
	return &AppConfig{
		Database: databaseConfig,
		Server:   serverConfig,
		JWT:      jWTConfig,
	}, nil
}
