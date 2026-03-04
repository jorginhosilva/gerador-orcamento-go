package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

func TestarPDF() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Teste do Gerador de Orcamento!")
    
    // Tenta salvar o arquivo na pasta atual
    err := pdf.OutputFileAndClose("teste.pdf")
    if err != nil {
        panic(err) // Se der erro (ex: pasta protegida), o programa avisa
    }
}