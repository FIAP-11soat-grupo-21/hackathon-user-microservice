package dto

import (
	"reflect"
	"testing"
)

// testDTOWithFields testa se um DTO foi criado com os campos corretos
func testDTOWithFields(t *testing.T, dto interface{}, expectedFields map[string]string) {
	t.Helper()

	v := reflect.ValueOf(dto)
	for fieldName, expectedValue := range expectedFields {
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			t.Errorf("Field %s not found in struct", fieldName)
			continue
		}

		actualValue := field.String()
		if actualValue != expectedValue {
			t.Errorf("Expected %s to be '%s', got '%s'", fieldName, expectedValue, actualValue)
		}
	}
}

// testEmptyDTO testa se todos os campos de um DTO estão vazios
func testEmptyDTO(t *testing.T, dto interface{}) {
	t.Helper()

	v := reflect.ValueOf(dto)
	typeOfDto := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfDto.Field(i).Name
		field := v.Field(i)

		if field.Kind() == reflect.String && field.String() != "" {
			t.Errorf("Expected %s to be empty, got '%s'", fieldName, field.String())
		}
	}
}
