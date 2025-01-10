package iteracao

import()

const quantidadeRepeticoes = 5

func Repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes += caracter
	}
	return repeticoes
}