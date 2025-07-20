package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Mapa para armazenar autores únicos
	autoresUnicos := make(map[string]bool)

	// Cria coletor padrão
	c := colly.NewCollector()

	// Ao encontrar cada citação, pega o autor
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		autor := e.ChildText(".author")
		if autor != "" {
			autoresUnicos[autor] = true
		}
	})

	// Trata erro de requisição
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Erro durante o scraping:", err)
	})

	// Visita a página principal
	err := c.Visit("https://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}

	// Exibe os autores únicos
	fmt.Println("👤 Autores únicos encontrados:")
	for autor := range autoresUnicos {
		fmt.Println("-", autor)
	}
}
