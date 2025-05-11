package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/tools"
)

type Node struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name"`
	Children []*Node  `json:"children,omitempty"`
}

type PathResult struct {
	Recipes      string  `json:"recipes"`
	Time         float64 `json:"time"`
	NodesVisited int     `json:"nodesVisited"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()
	for {
		var request struct {
			Element string `json:"element"`
			Algo 	string `json:"algorithm"`
			Mode 	string `json:"mode"`
			MaxRecipes int `json:"maxRecipes"`
		}
		if err := ws.ReadJSON(&request); err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		log.Printf("Received message: %v\n", request)
		log.Printf("Element: %s\n", request.Element)
		log.Printf("Algorithm: %s\n", request.Algo)
		log.Printf("Mode: %s\n", request.Mode)
		log.Printf("Max Recipes: %d\n", request.MaxRecipes)
		// harusnya di sini kita ngehandle request terus manggil fungsi yang sesuai (DFS/BFS)
		var response PathResult
		
		switch request.Mode {
		case "single":
			if request.Algo == "BFS" {
				steps, path := tools.RunBFS(request.Element)
				response.Recipes = path
				response.NodesVisited = steps
			} else if request.Algo == "DFS" {
				steps, path := tools.RunDFS(request.Element)
				response.Recipes = path
				response.NodesVisited = steps
			}
		case "multiple":
			if request.Algo == "BFS" {
				steps, path := tools.RunBFS(request.Element)
				response.Recipes = path
				response.NodesVisited = steps
			} else if request.Algo == "DFS" {
				steps, path := tools.RunDFS(request.Element)
				response.Recipes = path
				response.NodesVisited = steps
			}
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			break
		}

		if err := ws.WriteMessage(websocket.TextMessage, jsonResponse); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}