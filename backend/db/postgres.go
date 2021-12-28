package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/seeis/sensor-app/backend/constants"
	"github.com/seeis/sensor-app/backend/logger"
)

var PostgreSQL *sql.DB

func InitializePostgreSQLConnection() {
	var err error

	dataSourceName := constants.PostgresDriverName + "://" +
		os.Getenv("POSTGRES_USER") + ":" +
		os.Getenv("POSTGRES_PASSWORD") + "@" +
		os.Getenv("POSTGRES_ADDRESS") + "/" +
		os.Getenv("POSTGRES_DB") +
		constants.PostgresSSLDisabled

	PostgreSQL, err = sql.Open(constants.PostgresDriverName, dataSourceName)

	if err != nil {
		panic(err)
	}

	err = PostgreSQL.Ping()

	if err != nil {
		logger.Error.Printf("Connection to PostgreSQL failed: %s\n", err)
	}

	logger.Info.Println(constants.ConnectedToPostgreSQLMsg)
}
