package console

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	t.Run("contains a default value when no value supplied", func(t *testing.T) {
		args := []string{}
		defaults := map[string]string{"key1": "value1", "key2": "value2"}

		actual := ParseFlags(args, defaults)
		expected := "value1"

		if actual["key1"] != expected {
			t.Errorf("expected: '%s', actual: '%s'", expected, actual)
		}
	})

	t.Run("contains user-supplied value when no default provided", func(t *testing.T) {
		args := []string{"-key3=value.3"}
		defaults := map[string]string{"key1": "value1", "key2": "value2"}

		actual := ParseFlags(args, defaults)
		expected := "value.3"

		if actual["key3"] != expected {
			t.Errorf("expected: '%s', actual: '%s'", expected, actual)
		}
	})

	t.Run("replaces default value with user-defined value", func(t *testing.T) {
		args := []string{"-key1=value3"}
		defaults := map[string]string{"key1": "value1", "key2": "value2"}

		actual := ParseFlags(args, defaults)
		expected := "value3"

		if actual["key1"] != expected {
			t.Errorf("expected: '%s', actual: '%s'", expected, actual)
		}
	})
}
