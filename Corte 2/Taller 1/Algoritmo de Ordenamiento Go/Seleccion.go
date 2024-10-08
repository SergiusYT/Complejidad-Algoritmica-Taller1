package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 1000

func selectionSort(arr []int) (comparisons int, swaps int) {
	for i := 0; i < N-1; i++ {
		max_idx := i
		for j := i + 1; j < N; j++ {
			comparisons++
			if arr[j] > arr[max_idx] { // Cambio aquí: arr[j] > arr[max_idx]
				max_idx = j
			}
		}

		// Swap
		if max_idx != i {
			arr[i], arr[max_idx] = arr[max_idx], arr[i]
			swaps++
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Medir el tiempo de ejecución en nanosegundos
	start := time.Now()

	arr := make([]int, N)
	for i := range arr {
		arr[i] = rand.Intn(2501) + 300000 // Números entre 300000 y 302500
	}

	// Mostrar números generados
	fmt.Println("Numeros generados:")
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Ordenar usando el algoritmo de selección
	comparisons, swaps := selectionSort(arr)

	// Calcular la duración
	duration := time.Since(start)

	// Mostrar números ordenados
	fmt.Println("Numeros ordenados:")
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	fmt.Printf("Comparaciones: %d\n", comparisons)
	fmt.Printf("Intercambios: %d\n", swaps)
	fmt.Println("Estable: No")
	fmt.Println("Insercion: No")
	// Mostrar el tiempo de ejecución en nanosegundos
	fmt.Printf("Tiempo de ejecución en nanosegundos: %d\n", duration.Nanoseconds())
}
