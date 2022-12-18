package connection

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	path = "./config/initial_script.sql"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@/checkpoint_2")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).WithError(err).Info("error trying to read initial script")
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
	}

	return db, nil
}
