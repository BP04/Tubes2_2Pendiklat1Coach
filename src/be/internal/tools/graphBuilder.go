package tools

import (
	"encoding/json"
	"log"
	"os"
)

type Element struct {
	Name    string     `json:"element"`
	Tier    int        `json:"tier"`
	Recipes [][]string `json:"recipes"`
}

var (
	TierMap   = make(map[string]int)
	NameToID  = make(map[string]int)
	IDToName  []string
	Frequency = make(map[string]int)
	AdjList   = make([][][2]int, 0)
	elements  []Element
)

func ParseJSON() {
	file, err := os.Open("data/elements.json")
	if err != nil {
		log.Fatalf("Failed to open elements.json: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&elements); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
}

func BuildGraph() {
	var counter int = 0
	for _, element := range elements {
		if _, exists := NameToID[element.Name]; !exists {
			NameToID[element.Name] = counter
			IDToName = append(IDToName, element.Name)
			counter++
		}
		TierMap[element.Name] = element.Tier
	}

	AdjList = make([][][2]int, len(NameToID))

	for _, element := range elements {
		product := element.Name
		productID := NameToID[product]

		for _, recipe := range element.Recipes {
			if len(recipe) != 2 {
				continue
			}
			left, right := recipe[0], recipe[1]

			leftID := NameToID[left]
			rightID := NameToID[right]

			AdjList[productID] = append(AdjList[productID], [2]int{leftID, rightID})
		}
	}

	// fmt.Println("Graph built successfully!")
	// fmt.Printf("Total nodes: %d\n", len(AdjList))
}
