package main

func main() {
	type Item struct {
		ID			int // Identificador (ex: item 1, item 2, item 3)
		Nome 		string // Nome da planta
		Quantidade 	int // Quantidade do produto
		Preco 		float64 // Preço unitário do produto
		Total 		float64 // Total pedido
		
	}

	type Orcamento struct {
		Cliente 		string
		Data 			string
		Itens 			[]Item // slice da struct criada acima
		ValorTotal 		float64
	}
}