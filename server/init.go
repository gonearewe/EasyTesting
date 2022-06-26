package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Viper, the config loader, must be initialized first.
func initViper() {
	viper.SetConfigName("server-config") // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Config file found and successfully parsed
}

func initLogWriter() io.Writer {
	var file io.Writer
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	logPath := filepath.Join(exPath, viper.GetString("log_file_name"))
	file, err = os.Create(logPath)
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout) // write log to both stdout and file
	return writer
}

func initZeroLog(logWriter io.Writer) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = zerolog.New(logWriter).With().Timestamp().Caller().Logger()
	log.Info().Msg("ZeroLog initialized")
}

func initGin(logWriter io.Writer) {
	if !viper.GetBool("enable_console_color") {
		gin.DisableConsoleColor()
	}
	gin.DefaultWriter = logWriter
	gin.SetMode(gin.ReleaseMode)
	log.Info().Msg("Gin initialized")
}
