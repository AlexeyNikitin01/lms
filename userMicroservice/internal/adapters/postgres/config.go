package postgres

type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLmode  string
}
