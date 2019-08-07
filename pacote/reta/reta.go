package main

import "math"

// iniciando com letra MAISCULA é Publico (visível dentro e fora do pacote)
// iniciando com letra minuscula, ou _, é Privado (vísivel somente dentro do pacote)

// Ponto representa uma coordenada de um plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
