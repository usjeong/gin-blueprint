package model

import (
	"github.com/jmoiron/sqlx"
)

// DBPool datbase connection pool
type DBPool struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

// Example 데이터베이스 접근 예제
func (db *DBPool) Example(id int64) error {
	query := "SELECT id FROM test WHERE id = ?"
	_, err := db.Master.Query(query, id)
	return err
}
