package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (models *ApplicationShortApplicantModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (models *ApplicationLabelModels) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, models)
	case string:
		return json.Unmarshal([]byte(v), models)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (model *ApplicationApplicantModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (model *ApplicationSchoolWithDegreeModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (model *ApplicationAssignedUserModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (model *ApplicationStatusModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (model *ApplicationFileEakteModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (model *ApplicationFileFileModel) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
