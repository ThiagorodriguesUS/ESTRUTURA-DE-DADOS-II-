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

func minimo(no *No) *No {
	atual := no

	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}

	return atual
}

func Remover(no *No, v int) *No {
	if no == nil {
		return nil
	}

	if v < no.Valor {
		no.Esquerdo = Remover(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = Remover(no.Direito, v)
	} else {
		if no.Esquerdo == nil {
			return no.Direito
		}

		if no.Direito == nil {
			return no.Esquerdo
		}

		sucessor := minimo(no.Direito)
		no.Valor = sucessor.Valor
		no.Direito = Remover(no.Direito, sucessor.Valor)
	}

	return no
}

func (a *Arvore) Remover(v int) {
	a.Raiz = Remover(a.Raiz, v)
}

func EmOrdem(no *No) {
	if no != nil {
		EmOrdem(no.Esquerdo)
		fmt.Print(no.Valor, " ")
		EmOrdem(no.Direito)
	}
}

func criarArvore() *Arvore {
	arvore := &Arvore{}

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	return arvore
}

func main() {
	arvore1 := criarArvore()
	fmt.Println("Removendo folha 20:")
	arvore1.Remover(20)
	EmOrdem(arvore1.Raiz)
	fmt.Println()

	arvore2 := criarArvore()
	fmt.Println("Removendo nó 40 com um filho:")
	arvore2.Remover(40)
	EmOrdem(arvore2.Raiz)
	fmt.Println()

	arvore3 := criarArvore()
	fmt.Println("Removendo nó 50 com dois filhos:")
	arvore3.Remover(50)
	EmOrdem(arvore3.Raiz)
	fmt.Println()
}
