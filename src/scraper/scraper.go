package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"github.com/PuerkitoBio/goquery"
)

// struktur data untuk menyimpan nama elemen beserta recipe penyusunnya
type Element struct {
	Name    string     `json:"element"`
	Tier    int        `json:"tier"`
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

// fungsi buat ngehapus elemen dari slice
func deleteElmt[T any](arr []T, idx int) []T {
	if idx < 0 || idx >= len(arr) {
		return arr
	}
	return append(arr[:idx], arr[idx+1:]...)
}

// fungsi untuk ngehapus elemen sesuai di parameter fungsinya
func deleteElmtByName(elmts []Element, name string) []Element {
	for i, elmt := range elmts {
		if elmt.Name == name {
			return deleteElmt(elmts, i)
		}
	}
	return elmts
}

// fungsi untuk ngehapus recipe yang melibatkan bahan dengan nama yang ada di parameter fungsinya
func deleteRecipeByName(elmts []Element, name string) []Element {
	for i, elmt := range elmts {
		for j := len(elmt.Recipes) - 1; j >= 0; j-- {
			recipe := elmt.Recipes[j]
			if recipe[0] == name || recipe[1] == name {
				elmt.Recipes = deleteElmt(elmt.Recipes, j)
				elmts[i] = elmt
			}
		}
	}
	return elmts
}

func main() {
	os.Mkdir("icons", 0755)

	// get request ke halaman wiki Little Alchemy 2
	res, err := http.Get("https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// parsing HTML-nya
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var elements []Element
	var wg sync.WaitGroup
	var mu sync.Mutex

	inElements := make(map[string]bool)

	var currentTier int = 0

	// ambil semua elemen dari tabel
	doc.Find("h3, table.list-table").Each(func(i int, s *goquery.Selection) {
		// menentukan tier
		if goquery.NodeName(s) == "h3" {
			header := strings.ToLower(s.Text())
			if strings.Contains(header, "tier") {
				// pecah berdasarkan spasi
				// elemen kedua menyatakan tier
				parts := strings.Fields(header)
				if len(parts) >= 2 {
					if tierNum, err := strconv.Atoi(parts[1]); err == nil {
						currentTier = tierNum
					}
				}
			}
		} else if goquery.NodeName(s) == "table" {
			s.Find("tr").Each(func(i int, row *goquery.Selection) {
				if i == 0 {
					return
				}

				// ambil nama elemen di kolom pertama
				element := strings.TrimSpace(row.Find("td:nth-child(1) a").Text())
				if element == "" {
					return
				}

				// cek apakah elemen sudah diproses
				mu.Lock()
				if inElements[element] {
					mu.Unlock()
					return
				}
				inElements[element] = true
				mu.Unlock()

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

				wg.Add(1)

				go func(element, svgURL string, recipes [][]string, tier int) {
					defer wg.Done()
					// donlot SVG
					svgPath := filepath.Join("icons", element+".svg")
					if err := donlotFile(svgURL, svgPath); err != nil {
						log.Printf("Failed to download SVG for %s: %v", element, err)
						return
					}

					// tambahin elemen ke slice
					mu.Lock()
					elements = append(elements, Element{
						Name:    element,
						Tier:    tier,
						Recipes: recipes,
					})
					mu.Unlock()
				}(element, svgURL, recipes, currentTier)
			})
		}
	})

	// tunggu semua goroutine selesai
	wg.Wait()

	// ngehapusin elemen yang gaperlu berdasarkan QnA https://docs.google.com/spreadsheets/d/1SVCNEBOYS0_eKShaHFIrx_5YVOg-V1uiBX-fAHpypxg 
	// klo ga salah: time, ruins, archeologist, dan elemen yg muncul sbg bahan recipe 
	// tapi gaada di kolom elements dari laman ini https://little-alchemy.fandom.com/wiki/Elements_(Myths_and_Monsters)
	
	// maaf yh ini hardcode dikit soalnya gaada cara lain aowkaowkaokw
	elements = deleteElmtByName(elements, "Time")
	elements = deleteElmtByName(elements, "Ruins")
	elements = deleteElmtByName(elements, "Archeologist")
	elements = deleteRecipeByName(elements, "Time")
	elements = deleteRecipeByName(elements, "Ruins")
	elements = deleteRecipeByName(elements, "Archeologist")
	
	// ngeparsing laman myths and monsters
	hapusin := []string{
		"Time",
		"Ruins",
		"Archeologist"}

	res, err = http.Get("https://little-alchemy.fandom.com/wiki/Elements_(Myths_and_Monsters)")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.list-table").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, row *goquery.Selection) {
			if i == 0 {
				return
			}
			element := strings.TrimSpace(row.Find("td:nth-child(1) a").Text())
			if element == "" {
				return
			} else {
				// append ke slice hapusin
				hapusin = append(hapusin, element)
			}
		})
	})

	for _, elmt := range hapusin {
		elements = deleteElmtByName(elements, elmt)
		elements = deleteRecipeByName(elements, elmt)
	}

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
