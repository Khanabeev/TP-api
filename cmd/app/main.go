package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Khanabeev/TP-api/internal/app/apiserver"
	"github.com/Khanabeev/TP-api/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Info(".env file doesn't exists, continue with environment variables")
	}
}

func main() {
	envApi := struct {
		BindAddr string `toml:"bind_addr"`
	}{}
	envDb := struct {
		Connection string `toml:"connection"`
		Host       string `toml:"host"`
		Port       string `toml:"port"`
		User       string `toml:"user"`
		Password   string `toml:"password"`
		DbName     string `toml:"db_name"`
		SslMode    string `toml:"ssl_mode"`
	}{}

	_, err := toml.DecodeFile("configs/env/apiserver.toml", &envApi)
	if err != nil {
		logger.Error("File 'configs/env/apiserver.toml' doesn't exist or contains error: " + err.Error())
		log.Fatal(err)
	}

	_, err = toml.DecodeFile("configs/env/database.toml", &envDb)
	if err != nil {
		logger.Error("File 'configs/env/database.toml' doesn't exist or contains error: " + err.Error())
		log.Fatal(err)
	}

	serverConfig := &apiserver.ServerConfig{
		BindAddr: getEnvVar(envApi.BindAddr),
	}

	databaseConfig := &apiserver.DatabaseConfig{
		Connection: getEnvVar(envDb.Connection),
		Host:       getEnvVar(envDb.Host),
		Port:       getEnvVar(envDb.Port),
		User:       getEnvVar(envDb.User),
		Password:   getEnvVar(envDb.Password),
		DbName:     getEnvVar(envDb.DbName),
		SslMode:    getEnvVar(envDb.SslMode),
	}

	s := apiserver.New(serverConfig, databaseConfig)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func getEnvVar(varName string) string {
	envVar, exists := os.LookupEnv(varName)
	if !exists {
		log.Fatalf("Environment variable %s doesn't exist", varName)
	}

	return envVar
}
