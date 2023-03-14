package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Table struct {
	Id          int
	Exercise    string
	Description string
	Reference   string
}

func ConnectToDataBase(tableName string) []Table {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return ParseDB(db, tableName)
}

func ParseDB(db *sql.DB, tableName string) []Table {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", tableName)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tableData []Table
	for rows.Next() {
		t := Table{}
		if err := rows.Scan(&t.Id, &t.Exercise, &t.Description, &t.Reference); err != nil {
			panic(err)
		}
		tableData = append(tableData, t)
	}
	return tableData
}
