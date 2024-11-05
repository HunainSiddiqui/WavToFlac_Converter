package tests

import (
	"net/http/httptest"
	"os"
	"task-golang/router"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketAudioStreaming(t *testing.T) {

	server := httptest.NewServer(router.SetupRouter())
	defer server.Close()

	wsURL := "ws" + server.URL[4:] + "/ws/123"
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	wavData, err := os.ReadFile("test.wav")
	if err != nil {
		t.Fatalf("Failed to read test WAV file: %v", err)
	}

	err = conn.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		t.Fatalf("Failed to send WAV data: %v", err)
	}

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	_, flacData, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read FLAC response: %v", err)
	}

	assert.NotNil(t, flacData)
	assert.Greater(t, len(flacData), 0, "FLAC data should not be empty")

	err = os.WriteFile("output_integration.flac", flacData, 0644)
	if err != nil {
		t.Fatalf("Failed to write FLAC output: %v", err)
	}
}
