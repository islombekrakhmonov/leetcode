package config

import (
	"os"

	"github.com/spf13/cast"
)


const (
	CollectionName = "mongo"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	Environment   string
	MongoHost     string
	MongoPort     int
	MongoDatabase string
	MongoPassword string
	MongoUser     string
	LogLevel      string
	Port          string
	Host		  string
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue

}

func Load() Config {

	cfg := Config{}

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))
	cfg.MongoHost = cast.ToString(getOrReturnDefaultValue("MONGO_HOST", "localhost"))
	cfg.MongoPort = cast.ToInt(getOrReturnDefaultValue("MONGO_PORT", 27017))
	cfg.MongoDatabase = cast.ToString(getOrReturnDefaultValue("MONGO_DATABASE", "product_service"))
	cfg.MongoUser = cast.ToString(getOrReturnDefaultValue("MONGO_USER", "dell"))
	cfg.MongoPassword = cast.ToString(getOrReturnDefaultValue("MONGO_PASSWORD", "password"))
	cfg.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	cfg.Port = cast.ToString(getOrReturnDefaultValue("PORT", ":8080"))
	cfg.Host = cast.ToString(getOrReturnDefaultValue("HOST", "localhost"))

	return cfg

}