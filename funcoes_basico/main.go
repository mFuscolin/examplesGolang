package main

import (
	"fmt"

	"github.com/mFuscolin/funcoes_basico/matematica"
)

func main() {

	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)

	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)

	resultado = matematica.Soma(2, 3)

	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Divisor(6, 2)

	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)

	fmt.Printf("o resultado da divisão foi: %d e o resto da divisão foi %d\r\n", resultado, resto)
}
