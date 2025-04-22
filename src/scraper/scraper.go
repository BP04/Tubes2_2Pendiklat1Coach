package main
import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
	)

type Element struct {
	Name    string     `json:"element"`
	Recipes [][]string `json:"recipes"`
}

func main() {
	res, err := http.Get("https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)")
	if (err != nil) {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var elements []Element
	doc.Find("table.list-table tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		element := strings.TrimSpace(row.Find("td:nth-child(1) a").Text())
		if element == "" {
			return
		}

		var recipes [][]string
		row.Find("td:nth-child(2) li").Each(func(j int, li *goquery.Selection) {
			recipe := strings.TrimSpace(li.Text())
			parts := strings.Split(recipe, " + ")
			if len(parts) == 2 {
				recipes = append(recipes, []string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])})
			}
		})

		exists := false
		for _, e := range elements {
			if e.Name == element {
				exists = true
				break
			}
			for _, r := range e.Recipes {
				for _, recipe := range recipes {
					if (r[0] == recipe[0] && r[1] == recipe[1]) || (r[0] == recipe[1] && r[1] == recipe[0]) {
						exists = true
						break
					}
				}
				if exists {
					break
				}
			}
			if exists {
				break
			}
		}

		if !exists && len(recipes) > 0 {
			elements = append(elements, Element{
				Name:    element,
				Recipes: recipes,
			})
		}
	})

	file, err := os.Create("elements.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file) 
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(elements); err != nil {
		log.Fatal(err)
	}
}