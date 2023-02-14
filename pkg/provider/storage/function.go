package storage

// Ref: https://gosamples.dev/sqlite-intro/


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Func struct {
	Id         string
}

func InsertFunc(db *sql.DB, f *Func) error {
	_, err := db.Exec("INSERT INTO funcs(id) values(?)", f.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllFunc(db *sql.DB) ([] *Func, error) {
	rows, err := db.Query("SELECT * FROM funcs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var funcs []*Func
	for rows.Next() {
		var f Func
		if err := rows.Scan(&f.Id); err != nil {
			return nil, err
		}
		funcs = append(funcs, &f)
	}

	return funcs, nil
}

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./func.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS funcs (
		id TEXT PRIMARY KEY UNIQUE
	);
	`
	_, err = db.Exec(query)
	return db, err
}

func Cleanup(db *sql.DB) {
	db.Close()
	os.Remove("./func.db")
}