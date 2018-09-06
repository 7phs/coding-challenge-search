package common

import (
	"database/sql"
	"database/sql/driver"
	"time"
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

// NullInt represents a string that may be null, but more easy to use than sql.NullInt64.
// just set a value by pointer.
// NullInt implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt struct {
	V *int
}

// Scan implements the Scanner interface.
func (n *NullInt) Scan(value interface{}) (err error) {
	if n.V == nil {
		return
	}

	if value == nil {
		*n.V = 0

		return
	}

	v := sql.NullInt64{}
	if err := v.Scan(value); err != nil {
		return err
	}

	*n.V = int(v.Int64)
	return
}

// Value implements the driver Valuer interface.
func (n NullInt) Value() (driver.Value, error) {
	if n.V == nil {
		return nil, nil
	}

	return *n.V, nil
}

// NullInt64 represents a string that may be null, but more easy to use than sql.NullInt64.
// just set a value by pointer.
// NullInt64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt64 struct {
	V *int64
}

// Scan implements the Scanner interface.
func (n *NullInt64) Scan(value interface{}) (err error) {
	if n.V == nil {
		return
	}

	if value == nil {
		*n.V = 0

		return
	}

	v := sql.NullInt64{}
	if err := v.Scan(value); err != nil {
		return err
	}

	*n.V = v.Int64
	return
}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
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

// NullTime represents a time.Time that may be null.
// NullTime implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullTime struct {
	V *time.Time
}

// Scan implements the Scanner interface.
func (n *NullTime) Scan(value interface{}) (err error) {
	if n.V == nil {
		return
	}

	if value == nil {
		*n.V = time.Time{}

		return
	}

	switch v := value.(type) {
	case time.Time:
		*n.V = v
	case string:
		*n.V, err = time.Parse(time.RFC3339, v)
	}

	return
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if n.V == nil {
		return nil, nil
	}

	return *n.V, nil
}
