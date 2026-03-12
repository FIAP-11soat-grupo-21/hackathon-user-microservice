package exception

import "testing"

// ExceptionTest representa um caso de teste para exceções
type ExceptionTest struct {
	Name        string
	Message     string
	ExpectedMsg string
}

// testExceptionMessage testa se uma exceção retorna a mensagem esperada
func testExceptionMessage(t *testing.T, err error, expectedMsg string) {
	t.Helper()

	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}
