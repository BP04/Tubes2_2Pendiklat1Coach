package main

import (
	"fmt"
	"time"

	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/tools"
)

func main() {
	tools.ParseJSON()
	tools.BuildGraph()

	for _, element := range tools.IDToName {
		start := time.Now()

		steps, _ := tools.RunDFS(element)

		elapsed := time.Since(start)
		fmt.Printf("element: %s | DFS - Steps: %d | Time: %v\n", element, steps, elapsed)
	}
}
