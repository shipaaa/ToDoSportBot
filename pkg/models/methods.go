package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Table struct {
	Id          int
	Exercise    string
	Description string
	Reference   string
}

// Commit: Добавлена функции инициализации database функция вывода упражнений

func InitDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func GetAllExercisesFromDB(db *sql.DB, tableName string) []Table {
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
