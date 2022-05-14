package main

import "fmt"

func main() {
	entradas := [4][3]float64{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	pesos := [2]float64{0, 0}

	fmt.Printf("Pesos iniciais %v\n", pesos)
	fmt.Print("\n\n\n")

	taxaDeAprendizado := 0.5

	for ciclo := 0; ciclo < 5; ciclo++ {
		fmt.Printf("Ciclo %v\n", ciclo+1)

		for i, entrada := range entradas {
			fmt.Printf("Entrada %v\n", i+1)

			saida := calcSaida(entrada, pesos)

			fmt.Printf("Saida %v\n", saida)

			if saida != entrada[2] {
				pesos[0] = ajustarPeso(pesos[0], taxaDeAprendizado, entrada[2], saida, entrada[0])
				pesos[1] = ajustarPeso(pesos[1], taxaDeAprendizado, entrada[2], saida, entrada[1])
				fmt.Println("Saída diferente do esperado, pesos ajustados")
			}

			fmt.Printf("Pesos %v\n", pesos)
			fmt.Print("\n")
		}

		fmt.Print("\n\n\n\n\n")
	}
}

func calcSaida(entrada [3]float64, pesos [2]float64) float64 {
	return (entrada[0] * pesos[0]) + (entrada[1] * pesos[1])
}

func ajustarPeso(
	pesoAtual float64, taxaDeAprendizado float64, saidaEsperada float64, saida float64, entrada float64,
) float64 {
	return pesoAtual + (taxaDeAprendizado * (saidaEsperada - saida) * entrada)
}
