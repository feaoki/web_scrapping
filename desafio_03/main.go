package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Cria coletor
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	// Quando encontrar uma citação
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quote := e.ChildText(".text")
		author := e.ChildText(".author")
		fmt.Printf("📜 %s — %s\n", quote, author)
	})

	// Quando encontrar o link para a próxima página
	c.OnHTML("li.next > a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		fullURL := e.Request.AbsoluteURL(nextPage)
		fmt.Println("➡️ Indo para:", fullURL)
		e.Request.Visit(fullURL)
	})

	// Tratamento de erros
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Erro:", err)
	})

	// Inicia a visita na primeira página
	err := c.Visit("https://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}
}
