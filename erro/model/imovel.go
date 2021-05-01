package model

import "errors"

//Imovel representa informações do imovel
type Imovel struct {
	X     int    `json:"coordenada_X"`
	Y     int    `json:"coordenada_Y"`
	Nome  string `json:"Nome"`
	valor int
}

/*ErrValorDeImovelInvalido é um erro que é apresentado quando
é atribuido a imóvel um valor muito baixo*/
var ErrValorDeImovelInvalido = errors.New("O Valor  informado não é valido para um Imovel")

/*ErrValorDeImovelMuitoAlto  é um erro que é apresentado quando  é atribuido  ao
imovel  um valor muito alto  fora do mercado
*/
var ErrValorDeImovelMuitoAlto = errors.New("O Valor  informado é muito alto")

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
