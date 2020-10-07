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
		Ip:         "127.0.0.1",
		PrefixBits: 25,
	}
	expected := "01111111.00000000.00000000.00000001"
	result := addr.ToBinary()
	if result != expected {
		t.Errorf("expected %s to yield %s in binary, got %v", addr.Ip, expected, result)
	}
}
