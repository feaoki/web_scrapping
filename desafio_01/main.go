package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("programa 01")

	// Cria um coletor padrão
	c := colly.NewCollector()

	// Executa quando encontra um elemento com classe "quote"
	c.OnHTML(".quote", ColetarCitacao)

	// Trata erros
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Erro:", err)
	})

	// Visita a URL alvo
	err := c.Visit("https://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}
}

func ColetarCitacao(e *colly.HTMLElement) {
	citacao := e.ChildText(".text") // Busca texto da citação
	fmt.Println("Citação:", citacao)
	fmt.Println("passou pela função ColetarCitacao")
}
