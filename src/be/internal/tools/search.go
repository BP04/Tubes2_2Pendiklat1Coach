package tools

import (
	"encoding/json"
	"fmt"
	"time"
)

const INF = 1000000000

var (
	DFSIndex     = 0
	Aux          = make([][]int, 0)
	RAux         = make([][]int, 0)
	NextHop      = make([]int, 0)
	Label        = make([]string, 0)
	NodesVisited = 0
	DP           = make([]int, 0)
	IDToNode     = make([]int, 0)
)

type Node struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name"`
	Children []*Node `json:"children,omitempty"`
}

type PathResult struct {
	Recipes       []*Node `json:"recipes"`
	Time          float64 `json:"time"`
	TimeFormatted string  `json:"timeFormatted"`
	NodesVisited  int     `json:"nodesVisited"`
}

type Queue struct {
	items [][2]int
	head  int
	tail  int
	size  int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		items: make([][2]int, capacity),
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *Queue) Enqueue(item [2]int) {
	if q.size == len(q.items) {
		q.resize()
	}

	q.items[q.tail] = item
	q.tail = (q.tail + 1) % len(q.items)
	q.size++
}

func (q *Queue) Dequeue() [2]int {
	if q.size == 0 {
		return [2]int{}
	}

	item := q.items[q.head]
	q.head = (q.head + 1) % len(q.items)
	q.size--
	return item
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) resize() {
	newItems := make([][2]int, len(q.items)*2)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[(q.head+i)%len(q.items)]
	}
	q.items = newItems
	q.head = 0
	q.tail = q.size
}

type IntQueue struct {
	items []int
	head  int
	tail  int
	size  int
}

func NewIntQueue(capacity int) *IntQueue {
	return &IntQueue{
		items: make([]int, capacity),
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *IntQueue) Enqueue(item int) {
	if q.size == len(q.items) {
		q.resize()
	}

	q.items[q.tail] = item
	q.tail = (q.tail + 1) % len(q.items)
	q.size++
}

func (q *IntQueue) Dequeue() int {
	if q.size == 0 {
		return -1
	}

	item := q.items[q.head]
	q.head = (q.head + 1) % len(q.items)
	q.size--
	return item
}

func (q *IntQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *IntQueue) resize() {
	newItems := make([]int, len(q.items)*2)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[(q.head+i)%len(q.items)]
	}
	q.items = newItems
	q.head = 0
	q.tail = q.size
}

func InitSearch(targetElem string) {
	DFSIndex = 0
	Aux = make([][]int, 1)
	RAux = make([][]int, 1)
	NextHop = make([]int, 1)
	Label = make([]string, 1)
	NodesVisited = 0

	Label[0] = targetElem

	elementCount := len(IDToName)
	DP = make([]int, elementCount)
	IDToNode = make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		DP[i] = INF
		IDToNode[i] = -1
	}
}

func DFS(current int, adjIndex int) int {
	NodesVisited++
	name := IDToName[adjIndex]

	IDToNode[adjIndex] = current

	if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
		DP[adjIndex] = 0
		return 0
	}

	if DP[adjIndex] != INF {
		return DP[adjIndex]
	}

	productTier := TierMap[name]

	for _, combination := range AdjList[adjIndex] {
		leftID, rightID := combination[0], combination[1]
		leftName, rightName := IDToName[leftID], IDToName[rightID]
		leftTier, rightTier := TierMap[leftName], TierMap[rightName]

		if leftTier >= productTier || rightTier >= productTier {
			continue
		}

		DFSIndex++
		Aux = append(Aux, []int{})
		NextHop = append(NextHop, 0)
		Label = append(Label, "!")
		temp := DFSIndex

		Aux[current] = append(Aux[current], temp)

		var countLeft, countRight int

		if DP[leftID] == INF {
			DFSIndex++
			Aux = append(Aux, []int{})
			NextHop = append(NextHop, 0)
			Label = append(Label, leftName)
			Aux[temp] = append(Aux[temp], DFSIndex)
			countLeft = DFS(DFSIndex, leftID)
		} else {
			Aux[temp] = append(Aux[temp], IDToNode[leftID])
			countLeft = DP[leftID]
		}

		if DP[rightID] == INF {
			DFSIndex++
			Aux = append(Aux, []int{})
			NextHop = append(NextHop, 0)
			Label = append(Label, rightName)
			Aux[temp] = append(Aux[temp], DFSIndex)
			countRight = DFS(DFSIndex, rightID)
		} else {
			Aux[temp] = append(Aux[temp], IDToNode[rightID])
			countRight = DP[rightID]
		}

		totalCost := countLeft + countRight + 1
		if totalCost < DP[adjIndex] {
			DP[adjIndex] = totalCost
			NextHop[current] = temp
		}
	}

	return DP[adjIndex]
}

func BFS(targetElem string) int {
	targetID := NameToID[targetElem]
	BFSIndex := 0

	deg := make([]int, 1)
	minimum := make([]int, 1)
	p := NewIntQueue(100)
	q := NewQueue(100)

	for i := 0; i < len(IDToName); i++ {
		IDToNode[i] = -1
	}

	q.Enqueue([2]int{0, targetID})
	minimum[0] = INF
	Label[0] = targetElem
	IDToNode[targetID] = 0

	for !q.IsEmpty() {
		NodesVisited++
		pair := q.Dequeue()
		current := pair[0]
		adjIndex := pair[1]

		name := IDToName[adjIndex]

		if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
			minimum[current] = 0
			p.Enqueue(current)
			continue
		}

		productTier := TierMap[name]

		for _, combination := range AdjList[adjIndex] {
			leftID, rightID := combination[0], combination[1]
			leftName, rightName := IDToName[leftID], IDToName[rightID]
			leftTier, rightTier := TierMap[leftName], TierMap[rightName]

			if leftTier >= productTier || rightTier >= productTier {
				continue
			}

			BFSIndex++
			Aux = append(Aux, []int{})
			RAux = append(RAux, []int{})
			NextHop = append(NextHop, 0)
			deg = append(deg, 2)
			minimum = append(minimum, 1)
			Label = append(Label, "!")
			temp := BFSIndex

			Aux[current] = append(Aux[current], temp)
			RAux[temp] = append(RAux[temp], current)
			deg[current]++

			if IDToNode[leftID] == -1 {
				BFSIndex++
				Aux = append(Aux, []int{})
				RAux = append(RAux, []int{})
				NextHop = append(NextHop, 0)
				deg = append(deg, 0)
				minimum = append(minimum, INF)
				Label = append(Label, leftName)
				IDToNode[leftID] = BFSIndex

				Aux[temp] = append(Aux[temp], BFSIndex)
				RAux[BFSIndex] = append(RAux[BFSIndex], temp)
				q.Enqueue([2]int{BFSIndex, leftID})
			} else {
				Aux[temp] = append(Aux[temp], IDToNode[leftID])
				RAux[IDToNode[leftID]] = append(RAux[IDToNode[leftID]], temp)
			}

			if IDToNode[rightID] == -1 {
				BFSIndex++
				Aux = append(Aux, []int{})
				RAux = append(RAux, []int{})
				NextHop = append(NextHop, 0)
				deg = append(deg, 0)
				minimum = append(minimum, INF)
				Label = append(Label, rightName)
				IDToNode[rightID] = BFSIndex

				Aux[temp] = append(Aux[temp], BFSIndex)
				RAux[BFSIndex] = append(RAux[BFSIndex], temp)
				q.Enqueue([2]int{BFSIndex, rightID})
			} else {
				Aux[temp] = append(Aux[temp], IDToNode[rightID])
				RAux[IDToNode[rightID]] = append(RAux[IDToNode[rightID]], temp)
			}
		}
	}

	for !p.IsEmpty() {
		current := p.Dequeue()
		moves := minimum[current]

		if current == 0 {
			continue
		}

		for _, elem := range RAux[current] {
			deg[elem]--

			if Label[elem] == "!" {
				minimum[elem] += moves
			} else {
				if minimum[elem] > moves {
					minimum[elem] = moves
					NextHop[elem] = current
				}
			}

			if deg[elem] == 0 {
				p.Enqueue(elem)
			}
		}
	}

	return minimum[0]
}

func DFSBuildRecipes(current int, need int) []string {
	NodesVisited++

	name := Label[current]
	result := make([]string, 0)

	if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
		result = append(result, "("+name+")")
		return result
	}

	if name != "!" {
		for _, candidate := range Aux[current] {
			possible := DFSBuildRecipes(candidate, need)
			for _, s := range possible {
				result = append(result, "("+name+s+")")
				need--
				if need <= 0 {
					break
				}
			}
			if need <= 0 {
				break
			}
		}
		return result
	}

	leftRecipes := DFSBuildRecipes(Aux[current][0], need)
	leftLen := len(leftRecipes)

	rightNeeded := (need + leftLen - 1) / leftLen
	if rightNeeded < 1 {
		rightNeeded = 1
	}
	rightRecipes := DFSBuildRecipes(Aux[current][1], rightNeeded)

	for _, s := range rightRecipes {
		for _, t := range leftRecipes {
			result = append(result, t+s)
			need--
			if need <= 0 {
				break
			}
		}
		if need <= 0 {
			break
		}
	}

	return result
}

func BFSBuildRecipes(startNode int, need int) []string {
	memo := make([][]string, len(Aux))
	for i := range memo {
		memo[i] = make([]string, 0)
	}

	queue := NewIntQueue(100)
	deg := make([]int, len(Aux))

	for i := 0; i < len(Aux); i++ {
		name := Label[i]
		if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
			queue.Enqueue(i)
		}
		deg[i] = len(Aux[i])
	}

	for !queue.IsEmpty() {
		NodesVisited++

		node := queue.Dequeue()
		name := Label[node]

		if name == "Fire" || name == "Water" || name == "Earth" || name == "Air" || name == "Time" {
			memo[node] = append(memo[node], "("+name+")")
			for _, nextNode := range RAux[node] {
				deg[nextNode]--
				if deg[nextNode] == 0 {
					queue.Enqueue(nextNode)
				}
			}
			continue
		}

		if name != "!" {
			recipes := make([]string, 0)
			for _, recipeNode := range Aux[node] {
				if len(memo[recipeNode]) == 0 {
					continue
				}

				remaining := need - len(recipes)
				if remaining <= 0 {
					break
				}

				copyCount := 0
				if remaining < len(memo[recipeNode]) {
					copyCount = remaining
				} else {
					copyCount = len(memo[recipeNode])
				}
				for i := 0; i < copyCount; i++ {
					recipes = append(recipes, "("+name+memo[recipeNode][i]+")")
				}
			}
			memo[node] = recipes

			for _, nextNode := range RAux[node] {
				deg[nextNode]--
				if deg[nextNode] == 0 {
					queue.Enqueue(nextNode)
				}
			}
			continue
		}

		combined := make([]string, 0)
		left := Aux[node][0]
		right := Aux[node][1]

		if len(memo[left]) > 0 && len(memo[right]) > 0 {
			for _, l := range memo[left] {
				if len(combined) >= need {
					break
				}

				remaining := need - len(combined)
				copyCount := min(remaining, len(memo[right]))

				for i := 0; i < copyCount; i++ {
					combined = append(combined, l+memo[right][i])
				}
			}
		}

		memo[node] = combined

		for _, nextNode := range RAux[node] {
			deg[nextNode]--
			if deg[nextNode] == 0 {
				queue.Enqueue(nextNode)
			}
		}
	}

	return memo[startNode]
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
		ID:   0,
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

	InitSearch(targetElem)

	targetID := NameToID[targetElem]
	steps := DFS(0, targetID)

	elapsedTime := time.Since(startTime)
	elapsedNano := elapsedTime.Nanoseconds()

	timeFormatted := fmt.Sprintf("%.4f µs", float64(elapsedNano)/1000)

	recipe := BuildRecipeTree(0)
	recipe[0] = SortTree(recipe[0])

	result := PathResult{
		Recipes:       recipe,
		Time:          elapsedTime.Seconds(),
		TimeFormatted: timeFormatted,
		NodesVisited:  NodesVisited,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return steps, "{\"error\": \"Failed to generate JSON\"}"
	}

	return steps, string(jsonData)
}

func RunDFSMultiple(targetElem string, need int) string {
	startTime := time.Now()

	InitSearch(targetElem)

	targetID := NameToID[targetElem]
	DFS(0, targetID)

	NodesVisited = 0
	recipeList := DFSBuildRecipes(0, need)

	recipes := make([]*Node, len(recipeList))
	var err error

	for i, recipe := range recipeList {
		recipes[i], err = ParseTree(recipe)
		recipes[i] = SortTree(recipes[i])
		if err != nil {
			return "{\"error\": \"Failed to parse recipe tree\"}"
		}
	}

	elapsedTime := time.Since(startTime)
	elapsedNano := elapsedTime.Nanoseconds()

	timeFormatted := fmt.Sprintf("%.4f µs", float64(elapsedNano)/1000)

	result := PathResult{
		Recipes:       recipes,
		Time:          elapsedTime.Seconds(),
		TimeFormatted: timeFormatted,
		NodesVisited:  NodesVisited,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return "{\"error\": \"Failed to generate JSON\"}"
	}

	return string(jsonData)
}

func RunBFS(targetElem string) (int, string) {
	startTime := time.Now()

	InitSearch(targetElem)

	steps := BFS(targetElem)

	elapsedTime := time.Since(startTime)
	elapsedNano := elapsedTime.Nanoseconds()

	timeFormatted := fmt.Sprintf("%.4f µs", float64(elapsedNano)/1000)

	recipe := BuildRecipeTree(0)
	recipe[0] = SortTree(recipe[0])

	result := PathResult{
		Recipes:       recipe,
		Time:          elapsedTime.Seconds(),
		TimeFormatted: timeFormatted,
		NodesVisited:  NodesVisited,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return steps, "{\"error\": \"Failed to generate JSON\"}"
	}

	return steps, string(jsonData)
}

func RunBFSMultiple(targetElem string, need int) string {
	startTime := time.Now()

	InitSearch(targetElem)

	BFS(targetElem)

	NodesVisited = 0
	recipeList := BFSBuildRecipes(0, need)

	recipes := make([]*Node, len(recipeList))
	var err error

	for i, recipe := range recipeList {
		fmt.Println(recipe)
		recipes[i], err = ParseTree(recipe)
		recipes[i] = SortTree(recipes[i])
		if err != nil {
			return "{\"error\": \"Failed to parse recipe tree\"}"
		}
	}

	elapsedTime := time.Since(startTime)
	elapsedNano := elapsedTime.Nanoseconds()

	timeFormatted := fmt.Sprintf("%.4f µs", float64(elapsedNano)/1000)

	result := PathResult{
		Recipes:       recipes,
		Time:          elapsedTime.Seconds(),
		TimeFormatted: timeFormatted,
		NodesVisited:  NodesVisited,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return "{\"error\": \"Failed to generate JSON\"}"
	}

	return string(jsonData)
}
