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

func buscar(no *No, v int) *No {
	if no == nil {
		return nil
	}

	if no.Valor == v {
		return no
	}

	if v < no.Valor {
		return buscar(no.Esquerdo, v)
	}

	return buscar(no.Direito, v)
}

func (a *Arvore) Buscar(v int) *No {
	return buscar(a.Raiz, v)
}

func (a *Arvore) BuscarIter(v int) *No {
	atual := a.Raiz

	for atual != nil {
		if atual.Valor == v {
			return atual
		}

		if v < atual.Valor {
			atual = atual.Esquerdo
		} else {
			atual = atual.Direito
		}
	}

	return nil
}

func main() {
	arvore := &Arvore{}

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	if arvore.Buscar(65) != nil {
		fmt.Println("65 encontrado")
	} else {
		fmt.Println("65 não encontrado")
	}

	if arvore.Buscar(45) != nil {
		fmt.Println("45 encontrado")
	} else {
		fmt.Println("45 não encontrado")
	}

	if arvore.BuscarIter(65) != nil {
		fmt.Println("65 encontrado na busca iterativa")
	} else {
		fmt.Println("65 não encontrado na busca iterativa")
	}

	if arvore.BuscarIter(45) != nil {
		fmt.Println("45 encontrado na busca iterativa")
	} else {
		fmt.Println("45 não encontrado na busca iterativa")
	}
}
