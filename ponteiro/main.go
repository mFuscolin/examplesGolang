package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)

	//&casa verifica o endereço de memoria
	fmt.Printf("Casa é %p - %+v\r\n", &casa, casa)

	chacara := Imovel{4, 6, "Chacara foda", 6000000}

	mudaImovel(&chacara)
	fmt.Printf("Chacara é %p - %+v\r\n", &chacara, chacara)

}
func mudaImovel(imovel *Imovel) {

	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
s