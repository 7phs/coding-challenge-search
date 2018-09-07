package common

import (
	"database/sql"
	"database/sql/driver"
)

// NullString represents a string that may be null, but more easy to use than sql.NullString.
// just set a value by pointer.
// NullString implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullString struct {
	V *string
}

// Scan implements the Scanner interface.
func (n *NullString) Scan(value interface{}) (err error) {
	if n.V == nil {
		return
	}

	if value == nil {
		*n.V = ""

		return
	}

	v := sql.NullString{}
	if err := v.Scan(value); err != nil {
		return err
	}

	*n.V = v.String
	return
}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if n.V == nil {
		return nil, nil
	}

	return *n.V, nil
}

// NullFloat64 represents a string that may be null, but more easy to use than sql.NullInt64.
// just set a value by pointer.
// NullFloat64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullFloat64 struct {
	V *float64
}

// Scan implements the Scanner interface.
func (n *NullFloat64) Scan(value interface{}) (err error) {
	if n.V == nil {
		return
	}

	if value == nil {
		*n.V = 0

		return
	}

	v := sql.NullFloat64{}
	if err := v.Scan(value); err != nil {
		return err
	}

	*n.V = v.Float64
	return
}

// Value implements the driver Valuer interface.
func (n NullFloat64) Value() (driver.Value, error) {
	if n.V == nil {
		return nil, nil
	}

	return *n.V, nil
}
