package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"

	"github.com/Khanabeev/TP-api/internal/app/domain"
	"github.com/Khanabeev/TP-api/internal/app/service"
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

	masterRepository := domain.NewMasterRepository(dbClient)

	mh := MasterHandler{Service: service.NewMasterService(masterRepository)}

	// === ROUTES ===
	// Master
	s.router.HandleFunc("/api/v1/masters/{master_id:[0-9]+}", mh.GetMasterById).Methods(http.MethodGet)
	// ==============
	logger.Info("Starting api server at address " + s.serverConfig.BindAddr)

	return http.ListenAndServe(s.serverConfig.BindAddr, s.router)
}

func (s APIServer) getDbClient() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.databaseConfig.Host,
		s.databaseConfig.Port,
		s.databaseConfig.User,
		s.databaseConfig.Password,
		s.databaseConfig.DbName,
		s.databaseConfig.SslMode,
	)

	dbClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := dbClient.Ping(); err != nil {
		return nil, err
	}

	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 5)
	dbClient.SetMaxOpenConns(25)
	dbClient.SetMaxIdleConns(25)

	return dbClient, nil
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
