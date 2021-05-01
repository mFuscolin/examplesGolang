package main

import (
	"fmt"

	"github.com/mFuscolin/mapas/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache[casa.Nome] = casa

	fmt.Println("A uma casa amarela no cache?")

	imovel, achou := cache["Casa Amarela"]

	if achou {
		fmt.Printf("O que Achei foi: %+v\r\n", imovel)
	}

	apt := model.Imovel{}
	apt.Nome = "Apt azul"
	apt.X = 19
	apt.Y = 26
	apt.SetValor(70000)

	cache[apt.Nome] = apt

	fmt.Println("Quantos itens há no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave [%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]

	if achou {
		delete(cache, imovel.Nome)
	}

	//outra forma de fazer, sem retornar o imovel
	// _, achou = cache["Casa Amarela"]

	// if achou {
	// 	delete(cache, "Casa Amarela")
	// }

	fmt.Println("Quantos itens há no cache?", len(cache))
}
