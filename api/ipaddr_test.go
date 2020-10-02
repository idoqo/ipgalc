package api

import "testing"

func TestToBin(t *testing.T) {
	expected := "00010000"
	num := 16
	result := toBin(num)

	if result != expected {
		t.Errorf("expected toBin(%d) to return %s, got %s", num, expected, result)
	}
}

func TestIpToBinary(t *testing.T) {
	addr := &IPAddr{
		Ip:         "194.146.135.85",
		PrefixBits: 25,
	}
	expected := "11"
	result := addr.ToBinary()
	if result != expected {
		t.Errorf("expected %s to yield %s in binary, got %v", addr.Ip, expected, result)
	}
}
