package value_object

import "testing"

// TestCase representa um caso de teste genérico para validação
type TestCase struct {
	Name    string
	Input   string
	WantErr bool
}

// RunValidationTests executa testes de validação usando uma função de validação genérica
func RunValidationTests(t *testing.T, testName string, tests []TestCase, validationFunc func(string) error) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := validationFunc(tt.Input)
			if (err != nil) != tt.WantErr {
				t.Errorf("%s error = %v, wantErr %v", testName, err, tt.WantErr)
			}
		})
	}
}

// ValueTestCase representa um caso de teste para verificar valores retornados
type ValueTestCase struct {
	Name     string
	Input    string
	Expected string
}

// RunValueTests executa testes de valor usando uma função que retorna um valor
func RunValueTests(t *testing.T, tests []ValueTestCase, valueFunc func(string) (string, error)) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := valueFunc(tt.Input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
				return
			}

			if result != tt.Expected {
				t.Errorf("Expected '%s', got '%s'", tt.Expected, result)
			}
		})
	}
}
