package model

//Imovel representa informações do imovel
type Imovel struct {
	X     int    `json:"coordenada_X"`
	Y     int    `json:"coordenada_Y"`
	Nome  string `json:"Nome"`
	valor int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
