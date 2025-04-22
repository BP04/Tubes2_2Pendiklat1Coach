package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

// struktur data untuk menyimpan nama elemen beserta recipe penyusunnya
type Element struct {
	Name    string     `json:"element"`
	Recipes [][]string `json:"recipes"`
}

// fungsi untuk donlot file, di sini dipake untuk donlot SVG
func donlotFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	os.Mkdir("icons", 0755)

	// get request ke halaman wiki Little Alchemy 2
	res, err := http.Get("https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)")
	if (err != nil) {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// parsing HTML-nya
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ambil semua elemen dari tabel
	var elements []Element
	doc.Find("table.list-table tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		// ambil nama elemen di kolom pertama
		element := strings.TrimSpace(row.Find("td:nth-child(1) a").Text())
		if element == "" {
			return
		}
		
		// ambil URL dari SVG di kolom pertama
		svgURL := ""
		imgTag := row.Find("td:nth-child(1) .icon-hover a").First()
		if imgHref, exists := imgTag.Attr("href"); exists {
			svgURL = imgHref
		}

		// ambil semua recipe dari kolom kedua
		var recipes [][]string
		row.Find("td:nth-child(2) li").Each(func(j int, li *goquery.Selection) {
			recipe := strings.TrimSpace(li.Text())
			parts := strings.Split(recipe, " + ")
			if len(parts) == 2 {
				recipes = append(recipes, []string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])})
			}
		})

		// cek apakah elemen udah ada (biar gak duplikat)
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

		// kalo belum ada, tambahin
		if !exists && len(recipes) > 0 {
			elements = append(elements, Element{
				Name:    element,
				Recipes: recipes,
			})
			
			// donlot SVG
			svgPath := filepath.Join("icons", element+".svg")
			if err := donlotFile(svgURL, svgPath); err != nil {
				log.Printf("Failed to download SVG for %s: %v", element, err)
				return
			}
		}
	})

	// simpen data ke file JSON
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