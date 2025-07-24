package config

import (
	"app/internal/interfaces" 
)

func InitLogger() *interfaces.SlogLogger {
	// Configure structured logging with JSON output
	loggerStore := interfaces.NewSlogLogger()
	if loggerStore == nil {
		panic("Failed to initialize logger")
	}   
	loggerStore.Info("Logger initialized successfully")
	return loggerStore
}