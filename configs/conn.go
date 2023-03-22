package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// mysql database konektor
func PostgresConn() *sql.DB {
	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Host":     Config("DB_ADDRESS"),
		"DB_Name":     Config("DB_NAME"),
	}

	openConnection := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Name"])

	// Open the connection
	db, err := sql.Open("postgres", openConnection)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("DATABASE Successfully connected!")
	// return the connection
	return db
}
