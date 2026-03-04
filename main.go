package main

import (
    "fmt"
    "gerador-orcamento/pdf" // Importando o seu pacote local
)

func main() {
    fmt.Println("Gerando PDF de teste....")
    
    // Agora sim o prefixo 'pdf' vai funcionar perfeitamente
    pdf.TestarPDF() 
    
    fmt.Println("PDF gerado com sucesso! Verifique o arquivo teste.pdf")
}