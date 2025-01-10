package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz ola para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "") 
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Hola, e o nome da pessoa' quando é passado espanhol no segundo parametro", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Bonjour, e nome da pessoa' quando é passado o francês no segundo parâmetro", func(t *testing.T) {
		resultado := Ola("Alice","francês")
		esperado := "Bonjour, Alice"
		verificaMensagemCorreta(t, resultado, esperado)
	})


}