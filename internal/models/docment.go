package models

import "database/sql"

type Service struct {
	Code      string
	ServiceId sql.NullString
	URL       sql.NullString
}

type Document struct {
	Code   string
	NameKK sql.NullString
	NameRu sql.NullString
	NameEn sql.NullString
}