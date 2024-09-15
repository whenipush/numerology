package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Отправка запроса
	fmt.Println(reflect.TypeOf(url(1, 1, 1)))

	resp, err := http.Get(url(30, 4, 2003))
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Извлечение заголовков
	doc.Find(".detail-rozbor-items .detail-nadpisy-left").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		text := s.Text()

		// Выводим текст рецензии
		fmt.Printf("Review %d: %s\n", i, text)
	})

	doc.Find(".barvy-obalka-vetsi").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		text := s.Text()

		// Выводим текст рецензии
		fmt.Printf("Review %d: %s\n", i, text)
	})
}

func url(day int, mounth int, year int) string {
	url := fmt.Sprintf("https://ru.astro-seek.com/vychislit-numerologiya/?narozeni_den=%d&narozeni_mesic=%d&narozeni_rok=%d", day, mounth, year)
	return url
}
