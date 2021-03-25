package apiserver

import (
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
	logger.Info("Starting api server at address " + s.serverConfig.BindAddr)
	return http.ListenAndServe(s.serverConfig.BindAddr, s.router)
}
