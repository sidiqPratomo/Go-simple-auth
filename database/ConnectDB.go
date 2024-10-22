package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
	"github.com/sirupsen/logrus"
)

func ConnectDB(config *config.Config, log *logrus.Logger) *sql.DB {
	db, err := sql.Open("mysql", config.DbDsn) 
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("error connecting to DB")
		return nil
	}

	if err = db.Ping(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("error connecting to DB")
		return nil
	}

	return db
}