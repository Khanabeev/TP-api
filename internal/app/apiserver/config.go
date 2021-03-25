package apiserver

// ServerConfig - структура для конфигурационного файла сервера /configs/apiserver.toml
type ServerConfig struct {
	BindAddr string
}

// DatabaseConfig - структура для конфигурационного файла базы данных /configs/database.toml
type DatabaseConfig struct {
	Connection string
	Host       string
	Port       string
	User       string
	Password   string
	DbName     string
	SslMode    string
}

// NewServerConfig - инициализация конфига сервера с базовыми значениями
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		BindAddr: ":4000",
	}
}

// NewDatabaseConfig - инициализация конфига базы
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}
