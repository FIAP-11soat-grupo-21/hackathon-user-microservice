package env

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv string
	API   struct {
		Port string
		Host string
		URL  string
	}
	PasswordSalt int
	Database     struct {
		RunMigrations bool
		Host          string
		Name          string
		Port          string
		Username      string
		Password      string
	}
	AWS struct {
		Region  string
		Cognito struct {
			UserPoolId string
		}
	}
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.Load()
	})
	return instance
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) Load() {
	dotEnvPath := ".env"
	_, err := os.Stat(dotEnvPath)

	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	c.GoEnv = getEnv("GO_ENV")

	// API
	c.API.Port = getEnvWithDefault("API_PORT", "8080")
	c.API.Host = getEnvWithDefault("API_HOST", "0.0.0.0")
	c.API.URL = c.API.Host + ":" + c.API.Port

	// Password
	rawSalt := getEnvWithDefault("PASSWORD_SALT", "12")

	i64, err := strconv.ParseInt(rawSalt, 10, 32)
	if err != nil {
		log.Fatalf("Invalid PASSWORD_SALT value: %v", err)
	}

	c.PasswordSalt = int(i64)

	// Database
	c.Database.RunMigrations = getEnvWithDefault("DB_RUN_MIGRATIONS", "false") == "true"
	c.Database.Host = getEnvWithDefault("DB_HOST", "localhost")
	c.Database.Name = getEnvWithDefault("DB_NAME", "postgres")
	c.Database.Port = getEnvWithDefault("DB_PORT", "5432")
	c.Database.Username = getEnvWithDefault("DB_USERNAME", "postgres")
	c.Database.Password = getEnvWithDefault("DB_PASSWORD", "12345678")

	// AWS
	c.AWS.Region = getEnvWithDefault("AWS_REGION", "us-east-2")

	// AWS Cognito
	c.AWS.Cognito.UserPoolId = getEnvWithDefault("AWS_COGNITO_USER_POOL_ID", "test-pool-id")
}

func (c *Config) IsProduction() bool {
	return c.GoEnv == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.GoEnv == "development"
}
