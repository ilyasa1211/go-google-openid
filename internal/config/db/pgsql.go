package db

import (
	"log"
	"os"
	"strconv"
)

type PgsqlConf struct {
	User            string
	Password        string
	Host            string
	Port            int
	Dbname          string
	SslCertPath     string
	SslKeyPath      string
	SslRootCertPath string
}

func NewPgsqlConf() *PgsqlConf {
	envPort := os.Getenv("PGSQL_PORT")
	port, err := strconv.Atoi(envPort)

	if err != nil {
		log.Fatalln("pgsql port is not a valid number", err)
	}

	return &PgsqlConf{
		os.Getenv("PGSQL_USER"),
		os.Getenv("PGSQL_PASSWORD"),
		os.Getenv("PGSQL_HOST"),
		port,
		os.Getenv("PGSQL_DBNAME"),
		os.Getenv("PGSQL_SSL_CERT_PATH"),
		os.Getenv("PGSQL_SSL_KEY_PATH"),
		os.Getenv("PGSQL_SSL_ROOT_CERT_PATH"),
	}
}
