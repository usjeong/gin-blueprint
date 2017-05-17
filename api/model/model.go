package model

import (
	"github.com/jmoiron/sqlx"
)

// DBPool database connection pool
type DBPool struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

// Example DBPool usage
func (db *DBPool) Example(id int64) error {
	query := "SELECT id FROM test WHERE id = ?"
	_, err := db.Master.Query(query, id)
	return err
}
