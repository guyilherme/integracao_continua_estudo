package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(10, 10)

	if total != 20 {
		t.Errorf("Resultado soma incorreto. Era esperado %d e veio %d", 20, total)
	}

}
