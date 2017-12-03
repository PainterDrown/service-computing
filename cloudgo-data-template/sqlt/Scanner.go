package sqlt

// SqltScanner .
type Scanner interface {
	Scan(dest ...interface{}) error
}
