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

	casa := Imovel{}

	apartamento := Imovel{17, 56, "Meu apartamento", 76000}

	fmt.Printf("O apartamento é %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     60,
	}

	fmt.Printf("A chacara é : %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.valor = 450000
	casa.X = 60
	casa.Y = 30

	fmt.Printf("A casa é %+v\r\n", casa)

	sitio := Imovel{
		Nome:  "Sitio do pica pau amarelo",
		X:     80,
		Y:     90,
		valor: 5000000,
	}

	fmt.Printf("O sitio é %+v\r\n", sitio)
}
