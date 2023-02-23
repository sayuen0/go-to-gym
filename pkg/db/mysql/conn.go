package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"time"
)

var _db *sql.DB

func newMySQL(c *config.Config) (*sql.DB, error) {
	loc, err := time.LoadLocation(c.DB.Location)
	if err != nil {
		return nil, err
	}

	cfg := mysql.Config{
		DBName:    c.DB.DBName,
		User:      c.DB.User,
		Passwd:    c.DB.Password,
		Addr:      c.DB.Addr,
		Net:       c.DB.Net,
		ParseTime: c.DB.ParseTime,
		Collation: c.DB.Collation,
		Loc:       loc,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	boil.SetDB(db)
	boil.DebugMode = c.DB.Debug

	return db, nil
}

// GetConnection return sql.DB instance (global in application)
func GetConnection(c *config.Config) *sql.DB {
	if _db == nil {
		db, err := newMySQL(c)
		if err != nil {
			log.Fatal(err)
		}

		_db = db
	}
	return _db
}
