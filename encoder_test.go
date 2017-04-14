package rubymarshal

import (
	"bytes"
	"os/exec"
	"testing"
)

func checkAgainstRuby(t *testing.T, val interface{}, expected string) {
	b, err := Encode(val)
	if err != nil {
		t.Fatalf("Encode() failed: %s", err)
	}

	cmd := exec.Command("ruby", "test.rb")
	cmd.Stdin = bytes.NewReader(b)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("Error checking Ruby: %s", err)
	}

	result := out.String()
	if result != expected {
		t.Errorf("Encoded %v (%T), Ruby saw %s, expected %q", val, val, result, expected)
	}
}

func TestEncodeNil(t *testing.T) {
	checkAgainstRuby(t, nil, "nil")
}

func TestEncodeFalse(t *testing.T) {
	checkAgainstRuby(t, true, "true")
	checkAgainstRuby(t, false, "false")
}

func TestEncodeUint(t *testing.T) {
	checkAgainstRuby(t, 0, "0")
	checkAgainstRuby(t, 0xDE, "222")
	checkAgainstRuby(t, 0xDEAD, "57005")
	checkAgainstRuby(t, 0xDEADBE, "14593470")
	checkAgainstRuby(t, 0x3DEADBEE, "1038801902")
}

func TestEncodeInt(t *testing.T) {
	checkAgainstRuby(t, -0xDE, "-222")
	checkAgainstRuby(t, -0xDEAD, "-57005")
	checkAgainstRuby(t, -0xDEADBE, "-14593470")
	checkAgainstRuby(t, -0x3DEADBEE, "-1038801902")
}

func TestSlices(t *testing.T) {
	checkAgainstRuby(t, []int{}, "[]")
	checkAgainstRuby(t, []int{123}, "[123]")
	checkAgainstRuby(t, []interface{}{123, true, false, nil}, "[123, true, false, nil]")
}
