package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/neelchoudhary/budgetwallet-api-server/config"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"

	// Required imports
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

// ConnectDB connect to postgresql db
func ConnectDB(dbConfig *config.DBConfig) *sql.DB {
	fmt.Println("Connecting to Postgresql DB... " + dbConfig.Host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	utils.LogIfFatalAndExit(err)

	err = db.Ping()
	utils.LogIfFatalAndExit(err)

	MigrateUp(dbConfig)

	fmt.Println("You connected to your database.")

	return db
}

// MigrateUp migrate up to most recent migration
// Used for development and continuous integration
func MigrateUp(dbConfig *config.DBConfig) {
	fmt.Println("Migrating Up...")
	sourceURL := "file://postgresql/migrations"
	// postgres://user:password@host:port/dbname?query
	localDbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Dbname)
	m, err := migrate.New(sourceURL, localDbURL)
	utils.LogIfFatalAndExit(err)

	if err := m.Up(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Finished Migrations. Up to Date.")
}
