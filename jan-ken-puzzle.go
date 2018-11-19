/*
	Algorítmos Avançados
	Trabalho 1 - Jan-Ken-Puzzle

	Bruno Delmonde - 10262818
	Óliver S Becker - 10284890

*/
package main

import (
	"fmt"
	"sort"
)

type coord struct {
	x, y int
}

type end struct {
	x, y, tipo int
}

type ByPos []end

func (a ByPos) Len() int      { return len(a) }
func (a ByPos) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPos) Less(i, j int) bool {	// Função de comparar para ordenar.
	if a[i].x != a[j].x {
		return a[i].x < a[j].x
	}

	if a[i].y != a[j].y {
		return a[i].y < a[j].y
	}

	return a[i].tipo < a[j].tipo
}

func comida(eater int) int {	// Retorna qual peça pode ser comida pela que foi passada como parâmetro
	switch eater {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 1
	default:
		return -1
	}
}

func busca(tab [][]int, control []coord, R, C int, results *[]end) {
	tam := len(control)

	if tam == 1 {
		*results = append(*results, end{control[0].x + 1, control[0].y + 1, tab[control[0].x][control[0].y]})
		return
	}
	var ctemp coord

	for i := range control {
		curX := control[i].x
		curY := control[i].y
		tipo := tab[curX][curY]
		comida := comida(tipo)

		// Come pra baixo
		if curX+1 < R && tab[curX+1][curY] != 0 && comida == tab[curX+1][curY] {
			ctemp = control[i]
			tab[curX+1][curY] = tab[curX][curY]
			tab[curX][curY] = 0
			control[i] = control[tam-1]
			control = control[:tam-1]

			busca(tab, control, R, C, results)

			control = control[:tam]
			control[i] = ctemp
			tab[curX][curY] = tab[curX+1][curY]
			tab[curX+1][curY] = comida
		}

		// Come pra cima
		if curX > 0 && tab[curX-1][curY] != 0 && comida == tab[curX-1][curY] {
			ctemp = control[i]
			tab[curX-1][curY] = tab[curX][curY]
			tab[curX][curY] = 0
			control[i] = control[tam-1]
			control = control[:tam-1]

			busca(tab, control, R, C, results)

			control = control[:tam]
			control[i] = ctemp
			tab[curX][curY] = tab[curX-1][curY]
			tab[curX-1][curY] = comida
		}

		// Come pra direita
		if curY+1 < C && tab[curX][curY+1] != 0 && comida == tab[curX][curY+1] {
			ctemp = control[i]
			tab[curX][curY+1] = tab[curX][curY]
			tab[curX][curY] = 0
			control[i] = control[tam-1]
			control = control[:tam-1]

			busca(tab, control, R, C, results)

			control = control[:tam]
			control[i] = ctemp
			tab[curX][curY] = tab[curX][curY+1]
			tab[curX][curY+1] = comida
		}

		// Come pra esquerda
		if curY > 0 && tab[curX][curY-1] != 0 && comida == tab[curX][curY-1] {
			ctemp = control[i]
			tab[curX][curY-1] = tab[curX][curY]
			tab[curX][curY] = 0
			control[i] = control[tam-1]
			control = control[:tam-1]

			busca(tab, control, R, C, results)

			control = control[:tam]
			control[i] = ctemp
			tab[curX][curY] = tab[curX][curY-1]
			tab[curX][curY-1] = comida
		}
	}

	return
}

func comparaEnd(a, b end) bool {	// Verifica se duas variáveis do tipo 'end' são iguais
	return a.x == b.x && a.y == b.y && a.tipo == b.tipo
}

// Formata o vetor de saída para ficar conforme o esperado (mostrando o total de saídas, a quantidade
// de saídas diferentes e estas ordenadas)
func fomataSaida(raw []end) (total int, diferente int, final []end) {
	total = len(raw)
	var j int

	for i := range raw {
		for j = 0; j < len(final); j++ {
			if comparaEnd(raw[i], final[j]) == true {
				break
			}
		}

		if j == len(final) {
			diferente++
			final = append(final, raw[i])
		}
	}

	return
}

func main() {
	var R, C int
	var control []coord
	results := make([]end, 0)
	fmt.Scanf("%d %d", &R, &C)

	tab := make([][]int, R)

	for i := 0; i < R; i++ {	// Leitura da matriz do tabuleiro
		tab[i] = make([]int, C)
		for j := 0; j < C; j++ {
			fmt.Scanf("%d", &tab[i][j])
			if tab[i][j] != 0 {
				control = append(control, coord{i, j})
			}
		}
	}

	busca(tab, control, R, C, &results)	// Algorítmo backtracking

	sort.Sort(ByPos(results))	// Ordena o vetor de saídas

	leng, dif, fim := fomataSaida(results)	// Formata para o jeito de impressão desejado

	fmt.Printf("%d\n%d\n", leng, dif)	// Imprime a saída
	for i := range fim {
		fmt.Printf("%d %d %d\n", fim[i].x, fim[i].y, fim[i].tipo)
	}
}
