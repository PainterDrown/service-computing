package sqlt

import (
	"database/sql"
)

// SQLTemplate .
type SQLTemplate struct {
	SQLExecutor
}

// NewSQLTemplate .
func NewSQLTemplate(exe SQLExecutor) SQLTemplate {
	return SQLTemplate{exe}
}

// Insert .
func (tpl *SQLTemplate) Insert(query string, args ...interface{}) (int, error) {
	res, err := tpl.Exec(query, args...)
	id, _ := res.LastInsertId()
	return int(id), err
}

// Delete .
func (tpl *SQLTemplate) Delete(query string, args ...interface{}) (int, error) {
	res, err := tpl.Exec(query)
	af, _ := res.RowsAffected()
	return int(af), err
}

// Update .
func (tpl *SQLTemplate) Update(query string, args ...interface{}) error {
	_, err := tpl.Exec(query)
	return err
}

// QueryForRow .
func (tpl *SQLTemplate) QueryForRow(query string, args ...interface{}) *sql.Row {
	row := tpl.QueryRow(query, args)
	return row
}

// QueryForRows .
func (tpl *SQLTemplate) QueryForRows(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := tpl.Query(query, args)
	return rows, err
}
