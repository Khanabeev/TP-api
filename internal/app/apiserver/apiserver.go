package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Khanabeev/TP-api/pkg/logger"
	"github.com/gorilla/mux"
)

type APIServer struct {
	serverConfig   *ServerConfig
	databaseConfig *DatabaseConfig
	router         *mux.Router
}

func New(serverConfig *ServerConfig, databaseConfig *DatabaseConfig) *APIServer {
	return &APIServer{
		serverConfig:   serverConfig,
		databaseConfig: databaseConfig,
		router:         mux.NewRouter(),
	}
}

func (s APIServer) Start() error {
	dbClient, err := s.getDbClient()
	if err != nil {
		return err
	}

	logger.Info("Starting api server at address " + s.serverConfig.BindAddr)
	return http.ListenAndServe(s.serverConfig.BindAddr, s.router)
}

func (s APIServer) getDbClient() (*sql.DB, error) {

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.databaseConfig.Host,
		s.databaseConfig.Port,
		s.databaseConfig.User,
		s.databaseConfig.Password,
		s.databaseConfig.DbName,
		s.databaseConfig.SslMode,
	)
}
}
