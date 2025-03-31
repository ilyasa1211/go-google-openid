package migrations

import (
	"log"

	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/database/pgsql"
)

func CreateUserTable() {
	db := pgsql.NewPGSQLConn()

	if _, err := db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL)"); err != nil {
		log.Fatalln("error creating table user: ", err)
	}
}
