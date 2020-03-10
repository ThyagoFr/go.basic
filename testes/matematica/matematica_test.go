package matematica

import "testing"

const errorMessage = "Valor esperado %v, valor recebido %v."

func Testmedia(t *testing.T){
	valorEsperado := 7.28
	valor := media(7.2,9.9,6.1,5.9)
	if valor != valorEsperado {
		t.Errorf(errorMessage,valorEsperado,valor)
	}
}