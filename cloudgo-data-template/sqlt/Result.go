package sqlt

import (
	"database/sql"
)

// Result .
type Result struct {
	sql.Result
	*sql.Row
	*sql.Rows
}
