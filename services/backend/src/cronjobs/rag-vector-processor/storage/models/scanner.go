package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (models *DocumentCreatedFromModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *DocumentStatusModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
