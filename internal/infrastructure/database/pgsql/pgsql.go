package pgsql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ilyasa1211/go-google-openid/internal/config/db"
	_ "github.com/lib/pq"
)

func NewPGSQLConn() *sql.DB {
	conf := db.NewPgsqlConf()
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=verify-full sslcert=%s sslkey=%s sslrootcert=%s",
		conf.User, conf.Dbname, conf.Host, conf.Password, conf.SslCertPath, conf.SslKeyPath, conf.SslRootCertPath,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
