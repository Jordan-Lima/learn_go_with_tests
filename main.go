package main

import "fmt"

const (
	espanhol = "espanhol"
	frances = "francês"
	pfxOlaEs = "Hola, "
	pfxOlaFr = "Bonjour, "
	pfxOlaPt = "Olá, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "espanhol":
		prefixo = pfxOlaEs
	case "francês":
		prefixo = pfxOlaFr
	default:
		prefixo = pfxOlaPt
	}
	return
}

func main() {

	fmt.Println(Ola("mundo", ""))
}