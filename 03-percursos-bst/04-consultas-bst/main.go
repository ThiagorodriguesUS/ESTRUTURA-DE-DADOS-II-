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

func Altura(no *No) int {
	if no == nil {
		return -1
	}

	esquerda := Altura(no.Esquerdo)
	direita := Altura(no.Direito)

	if esquerda > direita {
		return esquerda + 1
	}

	return direita + 1
}

func Contar(no *No) int {
	if no == nil {
		return 0
	}

	return 1 + Contar(no.Esquerdo) + Contar(no.Direito)
}

func ContarFolhas(no *No) int {
	if no == nil {
		return 0
	}

	if no.Esquerdo == nil && no.Direito == nil {
		return 1
	}

	return ContarFolhas(no.Esquerdo) + ContarFolhas(no.Direito)
}

func Minimo(no *No) *No {
	if no == nil {
		return nil
	}

	for no.Esquerdo != nil {
		no = no.Esquerdo
	}

	return no
}

func Maximo(no *No) *No {
	if no == nil {
		return nil
	}

	for no.Direito != nil {
		no = no.Direito
	}

	return no
}

func main() {
	arvore := &Arvore{}

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	fmt.Println("Altura:", Altura(arvore.Raiz))
	fmt.Println("Quantidade de nós:", Contar(arvore.Raiz))
	fmt.Println("Quantidade de folhas:", ContarFolhas(arvore.Raiz))
	fmt.Println("Mínimo:", Minimo(arvore.Raiz).Valor)
	fmt.Println("Máximo:", Maximo(arvore.Raiz).Valor)

	vazia := &Arvore{}
	fmt.Println("Altura da árvore vazia:", Altura(vazia.Raiz))

	umNo := &Arvore{}
	umNo.Inserir(10)
	fmt.Println("Altura da árvore com um nó:", Altura(umNo.Raiz))
}
