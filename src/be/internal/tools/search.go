package tools

import (
	"encoding/json"
)

const INF = 1000000000

var (
	DFSIndex = 0
	Aux      = make([][]int, 0)
	NextHop  = make([]int, 0)
	Label    = make([]string, 0)
)

type Node struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

type PathResult struct {
	Steps int   `json:"steps"`
	Path  *Node `json:"path"`
}

func DFS(current int, adjIndex int) int {
	name := IDToName[adjIndex]

	if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
		return 0
	}

	minimum := INF
	productTier := TierMap[IDToName[adjIndex]]

	for _, combination := range AdjList[adjIndex] {
		leftID, rightID := combination[0], combination[1]
		leftTier := TierMap[IDToName[leftID]]
		rightTier := TierMap[IDToName[rightID]]

		if leftTier >= productTier || rightTier >= productTier {
			continue
		}

		DFSIndex++
		Aux = append(Aux, []int{})
		NextHop = append(NextHop, 0)
		Label = append(Label, "placeholder")
		temp := DFSIndex

		Aux[current] = append(Aux[current], temp)

		DFSIndex++
		Aux = append(Aux, []int{})
		NextHop = append(NextHop, 0)
		Label = append(Label, IDToName[leftID])
		Aux[temp] = append(Aux[temp], DFSIndex)
		countLeft := DFS(DFSIndex, leftID)

		DFSIndex++
		Aux = append(Aux, []int{})
		NextHop = append(NextHop, 0)
		Label = append(Label, IDToName[rightID])
		Aux[temp] = append(Aux[temp], DFSIndex)
		countRight := DFS(DFSIndex, rightID)

		if countLeft+countRight+1 < minimum {
			minimum = countLeft + countRight + 1
			NextHop[current] = temp
		}
	}

	return minimum
}

func TracebackJSON(current int) *Node {
	name := Label[current]

	if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
		return &Node{
			Type: "element",
			Name: name,
		}
	}

	if name == "placeholder" {
		return &Node{
			Type:  "combination",
			Left:  TracebackJSON(Aux[current][0]),
			Right: TracebackJSON(Aux[current][1]),
		}
	}

	return &Node{
		Type: "element",
		Name: name,
		Left: TracebackJSON(NextHop[current]),
	}
}

func RunDFS(targetElem string) (int, string) {
	DFSIndex = 0
	Aux = make([][]int, 1)
	NextHop = make([]int, 1)
	Label = make([]string, 1)

	targetID := NameToID[targetElem]
	Label[0] = targetElem

	steps := DFS(0, targetID)
	path := TracebackJSON(0)

	result := PathResult{
		Steps: steps,
		Path:  path,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return steps, "{\"error\": \"Failed to generate JSON\"}"
	}

	return steps, string(jsonData)
}
