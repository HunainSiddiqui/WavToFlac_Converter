package router

import (
	"net/http"
	handler "task-golang/controller"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
	
		return true
	},
}


func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	
	router.HandleFunc("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade WebSocket connection", http.StatusBadRequest)
			return
		}
		defer conn.Close()

		
		handler.WebSocketHandler(conn, mux.Vars(r)["id"])
	})

	return router
}
