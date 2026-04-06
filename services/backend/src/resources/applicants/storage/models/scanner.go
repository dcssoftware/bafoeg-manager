package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (models *ApplicantApplicantModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func UUIDToStringPtr(id uuid.UUID) sql.NullString {
	if id == uuid.Nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: id.String(), Valid: true}
}
