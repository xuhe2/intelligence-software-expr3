package initializers

import (
	"os"
	"testing"
)

// TestLoadEnv tests the LoadEnvVar function
func TestLoadEnv(t *testing.T) {
	LoadEnvVar("../.env")
	testEnvVar := os.Getenv("TEST_ENV_VAR")
	if testEnvVar != "ok" {
		t.Errorf("Expected 'ok', got '%s'", testEnvVar)
	}
}
