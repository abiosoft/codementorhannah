package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const mysqlSeed = `
CREATE TABLE IF NOT EXISTS` + "`user`" + `(
  ` + "`id`" + ` int(11) unsigned NOT NULL AUTO_INCREMENT,
  ` + "`username`" + `varchar(11) DEFAULT NULL,
  ` + "`password`" + `varchar(11) DEFAULT NULL,
  PRIMARY KEY (` + "`id`" + `)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
`

var mysqlHost, mysqlPort string

func init() {
	mysqlHost = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	mysqlPort = os.Getenv("MYSQL_PORT_3306_TCP_PORT")

	var db *sql.DB
	var err error
	for {
		time.Sleep(time.Second * 5)
		db, err = sql.Open("mysql", fmt.Sprintf("webapp:webapp_pass@tcp(%s:%s)/webapp_db", mysqlHost, mysqlPort))
		if err == nil {
			break
		} else {
			log.Println(err)
		}
	}

	_, err = db.Exec(mysqlSeed)
	if err != nil {
		log.Fatal(err)
	}

}
