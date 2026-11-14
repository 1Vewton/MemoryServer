package settings

import (
	"testing"
)

// Test the env reading of string var
func TestEnvReadingString(t *testing.T) {
	res := GetEnvString("test", "T")
	if res != "T" {
		t.Errorf("Expected T, got %s", res)
	}
	t.Setenv("test", "KK3S")
	res = GetEnvString("test", "T")
	if res != "KK3S" {
		t.Errorf("Expected KK3S, got %s", res)
	}
}

// TestEnvReadingInteger tests the env reading of integer var
func TestEnvReadingInteger(t *testing.T) {
	res := GetEnvInteger("TEST", 0)
	if res != 0 {
		t.Errorf(
			"Expected %d, got %d",
			0,
			res,
		)
	}
	t.Setenv("TEST", "1")
	res = GetEnvInteger("TEST", 0)
	if res != 1 {
		t.Errorf(
			"Expected %d, got %d",
			1,
			res,
		)
	}
}
