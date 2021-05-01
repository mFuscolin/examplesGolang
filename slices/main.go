package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Print(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Tokio"
	cidades[3] = "Singapura"
	capitais[1] = "Brasilia"
	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {

		if indice == 2 {
			cidade = "Joinville"
		}

		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}
}
