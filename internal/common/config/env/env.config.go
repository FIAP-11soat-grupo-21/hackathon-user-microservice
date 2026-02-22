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
	c.API.Port = getEnv("API_PORT")
	c.API.Host = getEnv("API_HOST")
	c.API.URL = c.API.Host + ":" + c.API.Port

	// Password
	rawSalt := getEnv("PASSWORD_SALT")

	i64, err := strconv.ParseInt(rawSalt, 10, 32)
	if err != nil {
		log.Fatalf("Invalid PASSWORD_SALT value: %v", err)
	}

	c.PasswordSalt = int(i64)

	// Database
	c.Database.RunMigrations = getEnv("DB_RUN_MIGRATIONS") == "true"
	c.Database.Host = getEnv("DB_HOST")
	c.Database.Name = getEnv("DB_NAME")
	c.Database.Port = getEnv("DB_PORT")
	c.Database.Username = getEnv("DB_USERNAME")
	c.Database.Password = getEnv("DB_PASSWORD")

	// AWS
	c.AWS.Region = getEnv("AWS_REGION")

	// AWS Cognito
	c.AWS.Cognito.UserPoolId = getEnv("AWS_COGNITO_USER_POOL_ID")
}

func (c *Config) IsProduction() bool {
	return c.GoEnv == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.GoEnv == "development"
}
