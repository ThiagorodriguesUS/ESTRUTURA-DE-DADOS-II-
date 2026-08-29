package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovaArvore() *Arvore {
	return &Arvore{}
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

func EmOrdem(no *No) {
	if no != nil {
		EmOrdem(no.Esquerdo)
		fmt.Print(no.Valor, " ")
		EmOrdem(no.Direito)
	}
}

func main() {
	arvore := NovaArvore()

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	fmt.Println("Valores da árvore:")
	EmOrdem(arvore.Raiz)
	fmt.Println()

	fmt.Println("Tentando inserir o valor 50 novamente:")
	arvore.Inserir(50)
	EmOrdem(arvore.Raiz)
	fmt.Println()

	fmt.Println("Teste com árvore inicialmente vazia:")

	outraArvore := NovaArvore()
	outraArvore.Inserir(10)

	EmOrdem(outraArvore.Raiz)
	fmt.Println()
}
