package main

import "fmt"

func main() {
	desafio1()
	divisao()
	desafio2()
}

// Aqui vou montar uam regra onde ele imprime de 1 a 100, e mostrar somente oss numero divisivel por 3
func desafio1() {
	for divitres := 1; divitres <= 100; divitres++ {
		if divitres%3 == 0 {
			fmt.Println(divitres)
		}
	}
}

//Aqui eu vou fazer uma divisao para separar os desafios
func divisao() {
	fmt.Println("--------------------------------")
}

//Aqui eu vou fazer ele mostra de 1 a 100, e toda vez que aparecer o numero divisivel por tres, ele mostra Pin e nos multiplos de 5 Pan e se for multiplo de 3 e 5 ele mostra PinPan

func desafio2() {
	for divitresecinco := 1; divitresecinco <= 100; divitresecinco++ {
		if divitresecinco%3 == 0 {
			fmt.Println("Pin =", divitresecinco)
		}
		if divitresecinco%5 == 0 {
			fmt.Println("Pan =", divitresecinco)
		}
		if divitresecinco%3 == 0 && divitresecinco%5 == 0 {
			fmt.Println("PinPan =", divitresecinco)
		}
	}
}
