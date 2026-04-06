package bafogapplication

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/pdf-form/bafog-application/model"
	"github.com/go-playground/validator/v10"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/form"
	pdfcpuModel "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ParseBAföGApplication(file []byte, validateStruct bool) (*model.RawBAföGApplicationBaseModel01, customerrors.ErrorInterface) {

	f := bytes.NewReader(file)

	conf := pdfcpuModel.NewDefaultConfiguration()
	conf.Cmd = pdfcpuModel.LISTFORMFIELDS

	ctx, err := api.ReadValidateAndOptimize(f, conf)
	if err != nil {
		panic(err)
	}

	fs, _, err := form.FormFields(ctx)
	if err != nil {
		panic(err)
	}

	applicationModel := &model.RawBAföGApplicationBaseModel01{}
	structValue := reflect.ValueOf(applicationModel).Elem()
	structTags := structValue.Type()

	for _, data := range fs {

		for structField := 0; structField < structTags.NumField(); structField++ {

			field := structTags.Field(structField)
			tag := field.Tag.Get("pdf")
			dateformat := field.Tag.Get("dateformat")
			emptyBoolTrue := field.Tag.Get("emptybooltrue")

			if dateformat == "" {
				dateformat = "02012006"
			}

			if data.Name == tag {
				fieldValue := structValue.Field(structField)

				if fieldValue.CanSet() {
					switch fieldValue.Type().String() {
					case "*string":
						value := ConvertText2PtrString(data.V)
						fieldValue.Set(reflect.ValueOf(value))
					case "string":
						fieldValue.SetString(strings.TrimSpace(data.V))
					case "bool":
						if emptyBoolTrue == "1" {
							fieldValue.SetBool(true)
							continue
						}
						fieldValue.SetBool(ConvertGerman2Bool(data.V))
					case "int", "int8", "int16", "int32", "int64":
						if v, err := strconv.ParseInt(data.V, 10, 64); err == nil {
							fieldValue.SetInt(v)
						}
					case "uint", "uint8", "uint16", "uint32", "uint64":
						if v, err := strconv.ParseUint(data.V, 10, 64); err == nil {
							fieldValue.SetUint(v)
						}
					case "float32", "float64":
						if v, err := strconv.ParseFloat(data.V, 64); err == nil {
							fieldValue.SetFloat(v)
						}
					case "model.DefaultFalseBoolean":
						if data.V == "" {
							fieldValue.SetBool(true)
						}
					case "*time.Time":
						date := ConvertText2PtrDate(data.V, dateformat)
						fieldValue.Set(reflect.ValueOf(date))

					case "*model.BescheidÜbermittelnAnEnum":
						value := ConvertText2PtrBescheidÜbermittelnAnEnum(data.V)
						fieldValue.Set(reflect.ValueOf(value))

					case "*model.VersichertEnum":
						value := ConvertText2PtrVersichertEnum(data.V)
						fieldValue.Set(reflect.ValueOf(value))
					case "*model.VerhältnisElternteileEnum":
						value := ConvertText2PtrVerhältnisElternteile(data.V)
						fieldValue.Set(reflect.ValueOf(value))
					case "*model.SorgeberechtigterEnum":
						value := ConvertText2PtrSorgeberechtigterEnum(data.V)
						fieldValue.Set(reflect.ValueOf(value))

					default:
						fmt.Println(fieldValue.Type().String())
						return nil, customerrors.NewInternalServerError(errors.New("cannot match variable"), "", "")
					}
				}
			}
		}
	}

	if validateStruct {
		validate := validator.New()
		validateErr := validate.Struct(applicationModel)
		if validateErr != nil {
			return nil, customerrors.NewValidationError(validateErr)
		}
	}

	return applicationModel, nil
}

func ConvertGerman2Bool(german string) bool {
	german = strings.ToLower(german)
	return german == "ja" || german == "yes"
}

func ConvertText2PtrString(input string) *string {
	trimmed := strings.TrimSpace(input)
	if trimmed != "" {
		return &trimmed
	}
	return nil
}

func ConvertText2PtrDate(input, dateformat string) *time.Time {
	if input == "" {
		return nil
	}

	if t, err := time.Parse(dateformat, input); err == nil {
		return &t
	}
	return nil
}

func ConvertText2PtrBescheidÜbermittelnAnEnum(input string) *model.BescheidÜbermittelnAnEnum {
	switch input {
	case "mich (stÃ¤ndiger Wohnsitz)", "mich (stÃ\u0083Â¤ndiger Wohnsitz)":
		result := model.BescheidÜbermittelnAnAntragstellerStändigerWohnsitz
		return &result
	case "mich (Wohnsitz am Ausbildungsort)":
		result := model.BescheidÜbermittelnAnAntragstellerAusbildungsortWohnsitz
		return &result
	case "meinen ersten Elternteil":
		result := model.BescheidÜbermittelnAnElternteil01
		return &result
	case "meinen zweiten Elternteil":
		result := model.BescheidÜbermittelnAnElternteil02
		return &result
	case "meine/-n Sorgeberechtigte/-n":
		result := model.BescheidÜbermittelnAnSorgeberechtiger
		return &result
	case "die von mir bevollmÃ¤chtigte Person":
		result := model.BescheidÜbermittelnAnBevollmächtigtePerson
		return &result
	default:
		result := model.BescheidÜbermittelnAnUnbekannt
		return &result
	}
}

func ConvertText2PtrVersichertEnum(input string) *model.VersichertEnum {
	switch input {
	case "gesetzlich familienversichert":
		result := model.VersichertGesetzlichFamilienversichert
		return &result
	case "studentisch familienversichert":
		result := model.VersichertStudentischFamilienversichert
		return &result
	case "privat versichert":
		result := model.VersichertPrivatVersichert
		return &result
	case "freiwillig gesetzlich versichert":
		result := model.VersichertFreiwilligGesetzlichVersichert
		return &result
	case "anders versichert":
		result := model.VersichertAndersVersichert
		return &result
	default:
		return nil
	}
}

func ConvertText2PtrVerhältnisElternteile(input string) *model.VerhältnisElternteileEnum {
	switch input {
	case "ja":
		result := model.VerhältnisElternteileLebenZusammen
		return &result
	case "ja, aber dauernd getrennt lebend":
		result := model.VerhältnisElternteileLebenZusammenAberDauerhaftGetrennt
		return &result
	case "nein":
		result := model.VerhältnisElternteileNichtZusammenlebend
		return &result
	default:
		return nil
	}
}

func ConvertText2PtrSorgeberechtigterEnum(input string) *model.SorgeberechtigterEnum {
	switch input {
	case model.SorgeberechtigterJa.String():
		result := model.SorgeberechtigterJa
		return &result
	case model.SorgeberechtigterNein.String():
		result := model.SorgeberechtigterNein
		return &result
	default:
		return nil
	}
}
