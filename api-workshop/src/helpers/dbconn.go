package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// DBConn Mysql Database Connection ...
func DBConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName))
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, errors.New("erro conexão db: " + err.Error())
	}
	return db, nil
}
