package main

import "fmt"

func main() {
	// Definindo o array
	array := []int{1, -4, 7, 12}

	// Chamando a função PositiveSum e imprimindo o resultado
	soma := PositiveSum(array)
	fmt.Printf("A soma dos números positivos é: %d\n", soma)
}

// Função para somar os números positivos em um array
func PositiveSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return sum
}
