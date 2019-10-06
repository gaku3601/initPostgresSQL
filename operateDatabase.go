package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

//OperateDatabase struct
type OperateDatabase struct {
	host         string
	database     string
	port         int
	user         string
	password     string
	sqlFilePaths []string
}

//NewOperateDatabase constractor
func NewOperateDatabase(host string, database string, port int, user string, password string, sqlFilePaths []string) *OperateDatabase {
	o := &OperateDatabase{host: host, database: database, port: port, user: user, password: password, sqlFilePaths: sqlFilePaths}
	o.allClearDatabase()
	o.allExecSQLFile()
	return o
}

func (o *OperateDatabase) allClearDatabase() {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", o.host, o.port, o.user, o.password, o.database))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("drop schema public cascade;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create schema public;")
	if err != nil {
		log.Fatal(err)
	}
}

func (o *OperateDatabase) allExecSQLFile() {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", o.host, o.port, o.user, o.password, o.database))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for _, file := range o.sqlFilePaths {
		fmt.Println(file)
		sql, err := readSQLFile(file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(sql)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err == nil {
		fmt.Println("complete🎉")
	}
}

func readSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	b := bytes.NewBuffer(content)
	return b.String(), nil
}
