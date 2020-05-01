package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/budgetmanagergrpc/config"

	"github.com/neelchoudhary/budgetmanagergrpc/utils"

	// Required imports
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	fmt.Println("You connected to your database.")

	return db
}
