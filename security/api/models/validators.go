package models

import (
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

type validateType struct {
	name  string
	param interface{}
}

func validates(value interface{}, fieldError *[]string, types ...validateType) {
	for _, d := range types {
		switch d.name {
		case "presence":
			validatePresenceOf(value, fieldError)
		case "maxLength":
			validateMaxLengthOf(value, fieldError, d.param.(int))
		case "minLength":
			validateMinLengthOf(value, fieldError, d.param.(int))
		}
	}
}

func validatePresenceOf(value interface{}, fieldError *[]string) {
	switch value.(type) {
	case string:
		if len(strings.TrimSpace(value.(string))) > 0 {
			return
		}
	case uint64:
		if value.(uint64) > 0 {
			return
		}
	}
	appendError(fieldError, "Preenchimento Obrigatório")
}

func validateMaxLengthOf(value interface{}, fieldError *[]string, length int) {
	switch value.(type) {
	case string:
		if len(strings.TrimSpace(value.(string))) <= length {
			return
		}
	}
	appendError(fieldError, fmt.Sprintf("Tamaho máximo permitido é de %d", length))
}

func validateMinLengthOf(value interface{}, fieldError *[]string, length int) {
	switch value.(type) {
	case string:
		if len(strings.TrimSpace(value.(string))) >= length {
			return
		}
	}
	appendError(fieldError, fmt.Sprintf("Tamaho minimo permitido é de %d", length))
}

func validateEmail(value string, fieldError *[]string) {
	if err := checkmail.ValidateFormat(value); err != nil {
		appendError(fieldError, "Email Inválido")
	}
}

func validatesForStringFields(
	required bool,
	minLength interface{},
	maxLength interface{},
) []validateType {
	var types []validateType
	if required {
		types = append(types, validateType{name: "presence"})
	}
	if maxLength != nil && maxLength.(int) > 0 {
		types = append(types, validateType{name: "maxLength", param: maxLength})
	}
	if minLength != nil && minLength.(int) > 0 {
		types = append(types, validateType{name: "minLength", param: minLength})
	}
	return types
}

func appendError(fieldError *[]string, message string) {
	*fieldError = append(*fieldError, message)
}
