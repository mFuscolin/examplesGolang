package main

import (
	"fmt"

	"github.com/mFuscolin/cursodego/pacotes/operadora"
	"github.com/mFuscolin/cursodego/pacotes/prefixo"
)

//NomeUsario nome do usuarios
var NomeUsuario = "murilo"

func main() {

	fmt.Printf("Nome do usuario %s\r\n", NomeUsuario)
	fmt.Printf("Nome da operadora %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo do Estado %s\r\n", prefixo.Numero)

}
