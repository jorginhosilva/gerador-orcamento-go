package models

import "fmt"

type Item struct {
	ID         int     // Identificador (ex: item 1, item 2, item 3)
	Nome       string  // Nome da planta
	Quantidade int     // Quantidade do produto
	Preco      float64 // Preço unitário do produto
	//Total      float64 // Total pedido

}

type Orcamento struct {
	Cliente    string
	Data       string
	Itens      []Item // slice da struct criada acima
	ValorTotal float64
}

func main() {
	item1 := Item{1, "Jaboticaba", 2, 250.00}
	item2 := Item{2, "Ponkan", 4, 100}

	meuOrcamento := Orcamento{
		Cliente: "André Luiz",
		Data:    "03/03/2026",
	}

	meuOrcamento.Itens = append(meuOrcamento.Itens, item1, item2)

	meuOrcamento.CalcularTotal()

	fmt.Printf("Orçamento de %s no valor de: R$ %.2f\n", meuOrcamento.Cliente, meuOrcamento.ValorTotal)

}

func (o *Orcamento) CalcularTotal() {
	var Total float64
	for _, item := range o.Itens {
		Total += item.Preco * float64(item.Quantidade)
	}

	o.ValorTotal = Total

}
