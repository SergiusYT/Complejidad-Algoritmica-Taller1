#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define N 1000

void insertionSort(int arr[], int n, int *comparisons, int *swaps) {
    int i, key, j;
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;

        // Comparar e insertar
        while (j >= 0 && arr[j] > key) {
            (*comparisons)++;
            arr[j + 1] = arr[j];
            j = j - 1;
            (*swaps)++;
        }
        arr[j + 1] = key;
        (*swaps)++;
    }
}

int main() {
    int arr[N];
    int comparisons = 0, swaps = 0;

    // Generar números aleatorios de 0 a 9
    srand(time(0));
    for (int i = 0; i < N; i++) {
        arr[i] = rand() % 10; // Números entre 0 y 9
    }

    // Mostrar números generados
    printf("Numeros generados:\n");
    for (int i = 0; i < N; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    // Ordenar usando el algoritmo de inserción
    insertionSort(arr, N, &comparisons, &swaps);

    // Mostrar números ordenados
    printf("Numeros ordenados:\n");
    for (int i = 0; i < N; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    // Mostrar resultados
    printf("Comparaciones: %d\n", comparisons);
    printf("Intercambios: %d\n", swaps);
    printf("Estable: Sí\n");
    printf("Insercion: Sí\n");

    return 0;
}
