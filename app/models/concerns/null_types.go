package concerns

import (
	"database/sql"
	"encoding/json"
)

// NullString type
type NullString struct {
	sql.NullString
}

// MarshalJSON handles how NullString should render to json
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(s.String)
}

// NullInt64 type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON handles how NullInt64 should render to json
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int64)
}

// NullFloat64 type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON handles how NullFloat64 should render to json
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Float64)
}
