package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Global logrus instance
var logg = logrus.New()

func init() {
	// Set log output to stdout
	logg.SetOutput(os.Stdout)

	// Set log format (customize as needed)
	logg.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		logg.SetLevel(logrus.DebugLevel)
	case "info":
		logg.SetLevel(logrus.InfoLevel)
	case "warn":
		logg.SetLevel(logrus.WarnLevel)
	case "error":
		logg.SetLevel(logrus.ErrorLevel)
	default:
		logg.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	// Load config from "conf/config.yml"
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Set log level from config
	setLogLevel(viper.GetString("log_level"))

	// Use logg globally
	logg.Info("Server starting...")
	logg.WithField("port", viper.GetInt("port")).Info("Listening on port")
}
