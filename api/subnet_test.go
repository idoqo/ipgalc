package api

import "testing"

func TestSplitIP(t *testing.T) {
	ip := "127.0.0.1"
	expected := []int{127, 0, 0, 1}

	result, err := splitIP(ip)
	if err != nil {
		t.Errorf("could not split %s: %w", ip, err)
	}
	for i := 0; i < 4; i++ {
		if result[i] != expected[i] {
			t.Errorf("expected result[%d] to be %d, got %d", i, expected[i], result[i])
		}
	}
}
