package tools

import (
	"encoding/json"
	"time"
)

const INF = 1000000000

var (
	DFSIndex    = 0
	Aux         = make([][]int, 0)
	RAux        = make([][]int, 0)
	NextHop     = make([]int, 0)
	Label       = make([]string, 0)
	NodesVisited = 0
)

type Node struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name"`
	Children []*Node  `json:"children,omitempty"`
}

type PathResult struct {
	Recipes      []*Node `json:"recipes"`
	Time         float64 `json:"time"`
	NodesVisited int     `json:"nodesVisited"`
}

// type Queue struct {
// 	items [][2]int
// 	head  int
// 	tail  int
// 	size  int
// }

func DFS(current int, adjIndex int) int {
	NodesVisited++
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
		Label = append(Label, "!")
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

func TracebackJSON(current int, recipeId int) *Node {
	name := Label[current]

	if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
		return &Node{
			Name: name,
		}
	}

	if name == "!" {
		left := TracebackJSON(Aux[current][0], 0)
		right := TracebackJSON(Aux[current][1], 0)
		
		children := []*Node{left, right}
		
		return &Node{
			Children: children,
		}
	}

	node := &Node{
		Name: name,
	}
	
	if NextHop[current] != 0 {
		recipe := TracebackJSON(NextHop[current], 0)
		node.Children = recipe.Children
	}
	
	return node
}

func BuildRecipeTree(current int) []*Node {
	recipes := []*Node{}
	
	mainRecipe := &Node{
		ID: 0,
		Name: Label[current],
	}
	
	if NextHop[current] != 0 {
		recipe := TracebackJSON(NextHop[current], 0)
		mainRecipe.Children = recipe.Children
	}
	
	recipes = append(recipes, mainRecipe)
	
	return recipes
}

func RunDFS(targetElem string) (int, string) {
	startTime := time.Now()
	
	DFSIndex = 0
	Aux = make([][]int, 1)
	NextHop = make([]int, 1)
	Label = make([]string, 1)
	NodesVisited = 0

	targetID := NameToID[targetElem]
	Label[0] = targetElem

	steps := DFS(0, targetID)
	recipes := BuildRecipeTree(0)

	result := PathResult{
		Recipes: recipes,
		Time: time.Since(startTime).Seconds(),
		NodesVisited: NodesVisited,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return steps, "{\"error\": \"Failed to generate JSON\"}"
	}

	return steps, string(jsonData)
}

// func NewQueue(capacity int) *Queue {
// 	return &Queue{
// 		items: make([][2]int, capacity),
// 		head:  0,
// 		tail:  0,
// 		size:  0,
// 	}
// }

// func (q *Queue) Enqueue(item [2]int) {
// 	if q.size == len(q.items) {
// 		q.resize()
// 	}

// 	q.items[q.tail] = item
// 	q.tail = (q.tail + 1) % len(q.items)
// 	q.size++
// }

// func (q *Queue) Dequeue() [2]int {
// 	if q.size == 0 {
// 		return [2]int{}
// 	}

// 	item := q.items[q.head]
// 	q.head = (q.head + 1) % len(q.items)
// 	q.size--
// 	return item
// }

// func (q *Queue) IsEmpty() bool {
// 	return q.size == 0
// }

// func (q *Queue) resize() {
// 	newItems := make([][2]int, len(q.items)*2)
// 	for i := 0; i < q.size; i++ {
// 		newItems[i] = q.items[(q.head+i)%len(q.items)]
// 	}
// 	q.items = newItems
// 	q.head = 0
// 	q.tail = q.size
// }

// type IntQueue struct {
// 	items []int
// 	head  int
// 	tail  int
// 	size  int
// }

// func NewIntQueue(capacity int) *IntQueue {
// 	return &IntQueue{
// 		items: make([]int, capacity),
// 		head:  0,
// 		tail:  0,
// 		size:  0,
// 	}
// }

// func (q *IntQueue) Enqueue(item int) {
// 	if q.size == len(q.items) {
// 		q.resize()
// 	}

// 	q.items[q.tail] = item
// 	q.tail = (q.tail + 1) % len(q.items)
// 	q.size++
// }

// func (q *IntQueue) Dequeue() int {
// 	if q.size == 0 {
// 		return -1
// 	}

// 	item := q.items[q.head]
// 	q.head = (q.head + 1) % len(q.items)
// 	q.size--
// 	return item
// }

// func (q *IntQueue) IsEmpty() bool {
// 	return q.size == 0
// }

// func (q *IntQueue) resize() {
// 	newItems := make([]int, len(q.items)*2)
// 	for i := 0; i < q.size; i++ {
// 		newItems[i] = q.items[(q.head+i)%len(q.items)]
// 	}
// 	q.items = newItems
// 	q.head = 0
// 	q.tail = q.size
// }

// func BFS(targetElem string) int {
// 	targetID := NameToID[targetElem]
// 	BFSIndex := 0

// 	deg := make([]int, 1)
// 	minimum := make([]int, 1)
// 	p := NewIntQueue(100)
// 	q := NewQueue(100)

// 	q.Enqueue([2]int{0, targetID})
// 	Aux = append(Aux, []int{})
// 	RAux = append(RAux, []int{})
// 	NextHop = append(NextHop, 0)
// 	minimum[0] = INF
// 	Label[0] = targetElem

// 	for !q.IsEmpty() {
// 		pair := q.Dequeue()
// 		current := pair[0]
// 		adjIndex := pair[1]

// 		name := IDToName[adjIndex]

// 		if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
// 			minimum[current] = 0
// 			p.Enqueue(current)
// 			continue
// 		}

// 		productTier := TierMap[name]

// 		for _, combination := range AdjList[adjIndex] {
// 			leftID, rightID := combination[0], combination[1]
// 			leftName, rightName := IDToName[leftID], IDToName[rightID]
// 			leftTier, rightTier := TierMap[leftName], TierMap[rightName]

// 			if leftTier >= productTier || rightTier >= productTier {
// 				continue
// 			}

// 			BFSIndex++
// 			Aux = append(Aux, []int{})
// 			RAux = append(RAux, []int{})
// 			NextHop = append(NextHop, 0)
// 			deg = append(deg, 2)
// 			minimum = append(minimum, 1)
// 			Label = append(Label, "!")
// 			temp := BFSIndex

// 			Aux[current] = append(Aux[current], temp)
// 			RAux[temp] = append(RAux[temp], current)
// 			deg[current]++

// 			BFSIndex++
// 			Aux = append(Aux, []int{})
// 			RAux = append(RAux, []int{})
// 			NextHop = append(NextHop, 0)
// 			deg = append(deg, 0)
// 			minimum = append(minimum, INF)
// 			Label = append(Label, leftName)

// 			Aux[temp] = append(Aux[temp], BFSIndex)
// 			RAux[BFSIndex] = append(RAux[BFSIndex], temp)
// 			q.Enqueue([2]int{BFSIndex, leftID})

// 			BFSIndex++
// 			Aux = append(Aux, []int{})
// 			RAux = append(RAux, []int{})
// 			NextHop = append(NextHop, 0)
// 			deg = append(deg, 0)
// 			minimum = append(minimum, INF)
// 			Label = append(Label, rightName)

// 			Aux[temp] = append(Aux[temp], BFSIndex)
// 			RAux[BFSIndex] = append(RAux[BFSIndex], temp)
// 			q.Enqueue([2]int{BFSIndex, rightID})
// 		}
// 	}

// 	for !p.IsEmpty() {
// 		current := p.Dequeue()
// 		moves := minimum[current]

// 		if current == 0 {
// 			continue
// 		}

// 		for _, elem := range RAux[current] {
// 			deg[elem]--

// 			if Label[elem] == "!" {
// 				minimum[elem] += moves
// 			} else {
// 				if minimum[elem] > moves {
// 					minimum[elem] = moves
// 					NextHop[elem] = current
// 				}
// 			}

// 			if deg[elem] == 0 {
// 				p.Enqueue(elem)
// 			}
// 		}
// 	}

// 	return minimum[0]
// }

// func RunBFS(targetElem string) (int, string) {
// 	Aux = make([][]int, 1)
// 	RAux = make([][]int, 1)
// 	NextHop = make([]int, 1)
// 	Label = make([]string, 1)

// 	steps := BFS(targetElem)
// 	path := TracebackJSON(0)

// 	result := PathResult{
// 		Steps: steps,
// 		Path:  path,
// 	}

// 	jsonData, err := json.Marshal(result)
// 	if err != nil {
// 		return steps, "{\"error\": \"Failed to generate JSON\"}"
// 	}

// 	return steps, string(jsonData)
// }
