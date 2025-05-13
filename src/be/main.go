package main

import (
	"log"
	"net/http"

	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/handler"
	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/scraper"
	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/tools"
)

func main() {
	// start := time.Now()

	// scraper.Scrape()

	// elapsed := time.Since(start)

	// fmt.Printf("Scraping took %s\n", elapsed)

	scraper.Scrape()
	tools.ParseJSON()
	tools.BuildGraph()

	http.HandleFunc("/ws", handler.WebSocketHandler)
	http.HandleFunc("/elements", handler.GetElements)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	// MaxStep := 0
	// MaxStepElement := ""
	// MaxTime := time.Duration(0)
	// MaxTimeElement := ""

	// for _, element := range tools.IDToName {
	// 	start := time.Now()

	// 	steps, _ := tools.RunBFS(element)

	// 	elapsed := time.Since(start)
	// 	fmt.Printf("element: %s | BFS - Steps: %d | Time: %v\n", element, steps, elapsed)

	// 	if elapsed > MaxTime {
	// 		MaxTime = elapsed
	// 		MaxTimeElement = element
	// 	}
	// 	if steps > MaxStep {
	// 		MaxStep = steps
	// 		MaxStepElement = element
	// 	}
	// }

	// fmt.Printf("element: %s | Time: %v\n", MaxTimeElement, MaxTime)
	// fmt.Printf("element: %s | Steps: %d\n", MaxStepElement, MaxStep)

	// steps, path := tools.RunBFS("Picnic")
	// fmt.Printf("Steps: %d\n", steps)
	// fmt.Printf("Path: %v\n", path)
}
