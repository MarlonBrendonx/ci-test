package main

import "testing"

func TestSum(t *testing.T) {
	var resp = Soma(10, 20)

	if resp != 30 {
		t.Errorf("Esperado %d - Recebido %d", 30, resp)
	}
}
