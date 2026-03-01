package utils

import "database/sql"

func NewNullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

func NewNullInt64(num int64) sql.NullInt64 {
	return sql.NullInt64{Int64: num, Valid: true}
}
