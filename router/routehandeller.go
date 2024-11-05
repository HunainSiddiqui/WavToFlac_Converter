package router

import (
	"net/http"
	handler "task-golang/controller"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader to upgrade HTTP requests to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Add custom origin check logic here if needed
		return true
	},
}

// SetupRouter initializes the Gorilla Mux router
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// WebSocket route
	router.HandleFunc("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP request to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade WebSocket connection", http.StatusBadRequest)
			return
		}
		defer conn.Close()

		// Call the WebSocket handler
		handler.WebSocketHandler(conn, mux.Vars(r)["id"])
	})

	return router
}
