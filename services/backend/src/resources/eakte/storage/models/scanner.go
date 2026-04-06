package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (models *EakteModelType) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *DokumentModelSource) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *DokumentModelFile) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *DokumentModelVorgang) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *DokumentModelEakte) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *EaktenApplicationMappingApplicationShortModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *EaktenApplicationMappingApplicantShortModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *EaktenApplicationMappingDegreeShortModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *EaktenApplicationMappingSchoolShortModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *EaktenApplicationMappingAkteShortModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
