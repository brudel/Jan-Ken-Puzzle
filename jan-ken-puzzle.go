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
	"strconv"
)

type coord struct {
	x, y int
}

type end struct {
	x, y, tipo int
}

type solucao struct {
	total, difs int
	resps []end
}

type ByPos []end

var mem map[string]int

var dfsVet = [4]int{1, -1, 0, 0}

func rdp(tab [][]int, results *solucao) bool {
	a, b := mem[tabToString(tab)]
	results.total += a
	return b
}

func temIlhas(x, y, p int, tab [][]int) int {
			//fmt.Printf("Ilha %d %d, valo %d, %d\n", x, y, tab[x][y], p)

	if !(tab[x][y] * p > 0) {
		//fmt.Printf("il Vorto\n")
		return 0
	}
	tab[x][y] = -tab[x][y]
		//fmt.Printf("il Válida\n")

	cont := 1

	if x < len(tab) - 1 {
		cont += temIlhas(x + 1, y, p, tab)
	}

	if x > 0 {
		cont += temIlhas(x - 1, y, p, tab)
	}

	if y < len(tab[0]) - 1 {
		cont += temIlhas(x, y + 1, p, tab)
	}

	if y > 0 {
		cont += temIlhas(x, y - 1, p, tab)
	}
		//fmt.Printf("ilCabo\n")

	return cont
}

func busca(tab [][]int, control []coord, R, C int, results *solucao) {

			fmt.Printf("Entro\n")
	tam := len(control)
	aux := 0
	if rdp(tab, results) {
			fmt.Printf("DP e vorto\n")
		return
	}

	if tam == 1 {
		results.resps = append(results.resps, end{control[0].x + 1, control[0].y + 1, tab[control[0].x][control[0].y]})
		results.total++
		results.difs++
		mem[tabToString(tab)] = 1
			fmt.Printf("Conto e vorto\n")
		return
	}

	if tam != temIlhas(control[0].x, control[0].y, 1, tab) {
		fmt.Printf("Ilha e vorto\n")
		temIlhas(control[0].x, control[0].y, -1, tab)
		mem[tabToString(tab)] = 0
		return
	}
	temIlhas(control[0].x, control[0].y, -1, tab)

	for i := range control {
		curX := control[i].x
		curY := control[i].y
		tipo := tab[curX][curY]
			fmt.Printf("t: %d\n", tipo)
		comida := tipo % 3 + 1
			fmt.Printf("c: %d\n", comida)

		// Come pra baixo
		if curX < R - 1 && tab[curX + 1][curY] != 0 && comida == tab[curX + 1][curY] {
			fmt.Printf("Baixo\n")

			tab[curX + 1][curY], tab[curX][curY], control[i], control[tam - 1], control = tab[curX][curY], 0, control[tam - 1], control[i], control[:tam - 1]

			busca(tab, control, R, C, results)
			aux += mem[tabToString(tab)]

			control = control[:tam]
			fmt.Printf("Tamanho de controle é %d, i: %d\n", len(control), i)
			control[i], control[i - 1] = control[i - 1], control[i]
			tab[curX][curY], tab[curX + 1][curY] = tab[curX + 1][curY], comida
		}

		// Come pra cima
		if curX > 0 && tab[curX - 1][curY] != 0 && comida == tab[curX - 1][curY] {
			fmt.Printf("Cima\n")

			tab[curX - 1][curY], tab[curX][curY], control[i], control[tam - 1], control = tab[curX][curY], 0, control[tam - 1], control[i], control[:tam - 1]

			busca(tab, control, R, C, results)
			aux += mem[tabToString(tab)]

			control = control[:tam]
			control[i], control[i - 1], tab[curX][curY], tab[curX - 1][curY] = control[i - 1], control[i], tab[curX - 1][curY], comida
		}

		// Come pra direita
		if curY+1 < C && tab[curX][curY+1] != 0 && comida == tab[curX][curY+1] {
			fmt.Printf("Dir\n")

			tab[curX][curY + 1], tab[curX][curY], control[i], control[tam - 1], control = tab[curX][curY], 0, control[tam - 1], control[i], control[:tam - 1]

			busca(tab, control, R, C, results)
			aux += mem[tabToString(tab)]

			control = control[:tam]
			control[i], control[i - 1], tab[curX][curY], tab[curX][curY + 1] = control[i - 1], control[i], tab[curX][curY + 1], comida
		}

		// Come pra esquerda
		if curY > 0 && tab[curX][curY-1] != 0 && comida == tab[curX][curY-1] {
			fmt.Printf("Esq\n")

			tab[curX][curY - 1], tab[curX][curY], control[i], control[tam - 1], control = tab[curX][curY], 0, control[tam - 1], control[i], control[:tam - 1]

			busca(tab, control, R, C, results)
			aux += mem[tabToString(tab)]

			control = control[:tam]
			control[i], control[i - 1], tab[curX][curY], tab[curX][curY - 1] = control[i - 1], control[i], tab[curX][curY - 1], comida
		}

		mem[tabToString(tab)] += aux
			fmt.Printf("Mudo\n")
	}
			fmt.Printf("Vorto\n")

	return
}

func (a ByPos) Len() int { //Len do Sort
	return len(a)
}

func (a ByPos) Swap(i, j int) { //Swap do Sort
	a[i], a[j] = a[j], a[i]
}

func (a ByPos) Less(i, j int) bool {	// Função de comparar para ordenar.
	if a[i].x != a[j].x {
		return a[i].x < a[j].x
	}

	if a[i].y != a[j].y {
		return a[i].y < a[j].y
	}

	return a[i].tipo < a[j].tipo
}

func main() {
	var R, C int
	var control []coord
	results := solucao{0, 0, make([]end, 0)}
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

	mem = make(map[string]int)

	busca(tab, control, R, C, &results)	// Algorítmo backtracking
	fmt.Printf("Fimm %d %d\n", results.total, len(results.resps))

	sort.Sort(ByPos(results.resps))	// Ordena o vetor de saídas

	fmt.Printf("%d\n%d\n", results.total, results.difs)	// Imprime a saída
	for i := range results.resps {
		fmt.Printf("%d %d %d\n", results.resps[i].x, results.resps[i].y, results.resps[i].tipo)
	}
}

/*
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
} */

/*
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

func comparaEnd(a, b end) bool {	// Verifica se duas variáveis do tipo 'end' são iguais
	return a.x == b.x && a.y == b.y && a.tipo == b.tipo
}

	return
} */

 func tabToString(tab [][]int) string {                                                              
         var key string;                                                                             
         key = ""                                                                                    
                                                                                                     
         for i := 0; i < len(tab); i++ {                                                                                                                                                       
                 for j := 0; j < len(tab[i]); j++ {                                                  
                         key += strconv.Itoa(tab[i][j])  // Converte o inteiro para string diretamente
                 }                                                                                   
         }                                                                                           
                                                                                                     
         return key                                                                                  
} 