package main

import (
	"fmt"

	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/database/pgsql/migrations"
)

func main() {
	migrations.CreateUserTable()

	fmt.Println("database migration succeed")
}
