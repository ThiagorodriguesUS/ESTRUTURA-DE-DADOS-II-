package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func inserir(no *No, v int) *No {
	if no == nil {
		return NovoNo(v)
	}

	if v < no.Valor {
		no.Esquerdo = inserir(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = inserir(no.Direito, v)
	}

	return no
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = inserir(a.Raiz, v)
}

func PreOrdem(no *No) {
	if no != nil {
		fmt.Print(no.Valor, " ")
		PreOrdem(no.Esquerdo)
		PreOrdem(no.Direito)
	}
}

func EmOrdem(no *No) {
	if no != nil {
		EmOrdem(no.Esquerdo)
		fmt.Print(no.Valor, " ")
		EmOrdem(no.Direito)
	}
}

func PosOrdem(no *No) {
	if no != nil {
		PosOrdem(no.Esquerdo)
		PosOrdem(no.Direito)
		fmt.Print(no.Valor, " ")
	}
}

func EmLargura(no *No) {
	if no == nil {
		return
	}

	fila := []*No{no}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		fmt.Print(atual.Valor, " ")

		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}

		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}
}

func main() {
	arvore := &Arvore{}

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	fmt.Print("Pré-ordem: ")
	PreOrdem(arvore.Raiz)

	fmt.Print("\nEm-ordem: ")
	EmOrdem(arvore.Raiz)

	fmt.Print("\nPós-ordem: ")
	PosOrdem(arvore.Raiz)

	fmt.Print("\nEm largura: ")
	EmLargura(arvore.Raiz)

	fmt.Println()
}
