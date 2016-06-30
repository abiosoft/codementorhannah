package basictesting

import (
	"testing"
)

/**
1. File name must be end with '_test.go'
2. Function name must start with 'Test' and must have definition (*testing.T)
3. Table driven tests - tests in array (tabular form), you'll validate via a loop.
**/

func TestSayHello(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Dave", "Hello Dave"},
		{"Ann", "Hello Ann"},
		{"hannah", "Hello hannah"},
	}

	var output string
	for i, test := range tests {
		output = sayHello(test.input)
		if output != test.expected {
			t.Errorf("Test %d: expected %s found %s", i, test.expected, output)
		}
	}
}

func TestSayHi(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Dave", "Hi Dave"},
		{"Ann", "Hi Ann"},
	}

	var output string
	for i, test := range tests {
		output = sayHi(test.input)
		if output != test.expected {
			t.Errorf("Test %d: expected %s found %s", i, test.expected, output)
		}
	}
}

func TestConvertToInt(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		shouldError bool
	}{
		{"Dave", 0, true},
		{"10", 10, false},
	}

	for i, test := range tests {
		output, err := convertToInt(test.input)
		// if an error is expected
		if test.shouldError && err == nil {
			t.Errorf("Test %d: error expected but no error", i)
		}
		if output != test.expected {
			t.Errorf("Test %d: expected %d found %d", i, test.expected, output)
		}
	}
}

func TestHello(t *testing.T) {
	h, err := NewHello()

	if _, ok := err.(MyError); !ok {
		t.Error("Error type MyError expected")
		// no point going forward, h is nil
	}

	h.Say() // h is still nil
}
