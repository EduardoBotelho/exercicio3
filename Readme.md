# Soma de Números Positivos em um Array (Golang)

## Este programa em Go percorre um array de números inteiros, filtra apenas os números positivos e calcula a soma total. É uma solução simples e eficiente para lidar com operações básicas de array em Go.

---

## 📋 Funcionalidades

# - Recebe um array de números inteiros.
# - Filtra os números positivos.
# - Retorna e exibe no terminal a soma dos números positivos.

---

## 💻 Código

## ```go
## package main

## import "fmt"

## func main() {
##	// Definindo o array
##	array := []int{1, -4, 7, 12}
	
##	// Chamando a função PositiveSum e imprimindo o resultado
##	soma := PositiveSum(array)
##	fmt.Printf("A soma dos números positivos é: %d\n", soma)
##}

## // Função para somar os números positivos em um array
## func PositiveSum(numbers []int) int {
##	sum := 0
##	for _, num := range numbers {
## 		if num > 0 {
##			sum += num
##		}
##	}
##	return sum
##}
#   e x e r c i c i o 3 
 
 
