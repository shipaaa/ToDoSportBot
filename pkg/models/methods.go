package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Table struct {
	Id          int
	Exercise    string
	Description string
	Reference   string
}

func InitDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetDataFromDB(db *sql.DB, query string) []Table {
	rows, err := db.Query(query)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()

	var tableData []Table
	for rows.Next() {
		t := Table{}
		if err = rows.Scan(&t.Id, &t.Exercise, &t.Description, &t.Reference); err != nil {
			log.Error(err)
		}
		tableData = append(tableData, t)
	}
	return tableData
}

func GenerateQuery(exercise string) string {
	return fmt.Sprintf("SELECT * FROM %s ORDER BY id", exercise)
}
