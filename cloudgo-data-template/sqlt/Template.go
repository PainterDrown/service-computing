package sqlt

// Template .
type Template struct {
	Executor
}

// NewTemplate .
func NewTemplate(exe Executor) Template {
	return Template{exe}
}

// Exec .
func (tpl *Template) Exec(sql string, cb Callback, args ...interface{}) error {
	res, err := tpl.Executor.Exec(sql, args...)
	cb(Result{res})
	return err
}

// Query .
func (tpl *Template) Query(sql string, args ...interface{}) (int, error) {
	res, err := tpl.Exec(sql)
	af, _ := res.RowsAffected()
	return int(af), err
}

// SqltQueryRow .
func (tpl *Template) QueryRow(sql string, args ...interface{}) error {
	_, err := tpl.Exec(sql)
	return err
}
