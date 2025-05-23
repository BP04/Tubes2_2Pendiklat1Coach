package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/BP04/Tubes2_2Pendiklat1Coach/internal/tools"
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

// fungsi untuk ngoper ./data/elements.json ke frontend,
func GetElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Content-Type", "application/json")
	file, err := os.Open("data/elements.json")
	if err != nil {
		http.Error(w, "Failed to open elements.json", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var elements []tools.Element
	if err := json.NewDecoder(file).Decode(&elements); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(elements); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
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

		// harusnya di sini kita ngehandle request terus manggil fungsi yang sesuai (DFS/BFS)
		var response PathResult

		var path string
		switch request.Mode {
		case "single":
			if request.Algo == "BFS" {
				_, path = tools.RunBFS(request.Element)
			} else if request.Algo == "DFS" {
				_, path = tools.RunDFS(request.Element)
			}
		case "multiple":
			if request.Algo == "BFS" {
				path = tools.RunBFSMultiple(request.Element, request.MaxRecipes)
			} else if request.Algo == "DFS" {
				path = tools.RunDFSMultiple(request.Element, request.MaxRecipes)
			}
		}
		
		if err := json.Unmarshal([]byte(path), &response); err != nil {
			log.Println("Error unmarshalling JSON:", err)
			break
		}

		// buat debug doang
		// prettyPath, err := json.MarshalIndent(response, "", "  ")
		// log.Println("Pretty JSON:", string(prettyPath))
		
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			break
		}

		// log.Println("Response JSON:", string(jsonResponse)) // buat debug doang

		if err := ws.WriteMessage(websocket.TextMessage, jsonResponse); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}