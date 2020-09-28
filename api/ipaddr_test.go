package api

import "testing"

func testToBin(t *testing.T) {
	expected := "00010000"
	num := 16
	result := toBin(num)

	if result != expected {
		t.Errorf("expected toBin(%d) to return %s, got %s", num, expected, result)
	}
}
