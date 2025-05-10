package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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

		// harusnya di sini kita ngehandle request terus manggil fungsi yang sesuai (DFS/BFS)
		// var response PathResult

		// yg ini cuma contoh bang buat ngetes
		response := PathResult{
			Recipes:      []*Node{
				{Name: "Land",
					Children: []*Node{
						{Name: "Earth"},
						{Name: "Earth"},
					},
				},
			},
			Time:         0,
			NodesVisited: 0,
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