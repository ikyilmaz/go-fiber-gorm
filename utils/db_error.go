package utils

type DBError struct {
	Message string
}

func NewDBError(err error) *DBError { return &DBError{err.Error()} }

func (d *DBError) Error() string { return d.Message }
