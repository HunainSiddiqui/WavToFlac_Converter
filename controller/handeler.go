package controller

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
)

const maxWorkers = 5

var (
	jobQueue   = make(chan *ConversionJob, maxWorkers)
	once       sync.Once
	workerPool []Worker
)

type ConversionJob struct {
	conn    *websocket.Conn
	id      string
	wavData []byte
}

type Worker struct {
	ID int
}

func NewWorker(id int) Worker {
	return Worker{ID: id}
}

func (w Worker) Start() {
	for job := range jobQueue {
		processJob(job)
	}
}

func InitializeWorkerPool() {
	workerPool = make([]Worker, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(i)
		go worker.Start()
		workerPool[i] = worker
	}
}

func WebSocketHandler(conn *websocket.Conn, id string) {

	once.Do(InitializeWorkerPool)

	log.Printf("Client connected with ID: %s", id)
	defer conn.Close()

	for {

		messageType, wavData, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		if messageType == websocket.BinaryMessage {

			job := &ConversionJob{
				conn:    conn,
				id:      id,
				wavData: wavData,
			}
			jobQueue <- job
		} else {
			log.Println("Unsupported message type:", messageType)
		}
	}
}

func processJob(job *ConversionJob) {

	flacData, err := ConvertWAVToFLAC(job.wavData)
	if err != nil {
		log.Println("Conversion error:", err)
		errMsg := fmt.Sprintf("Conversion error: %v", err)
		job.conn.WriteMessage(websocket.TextMessage, []byte(errMsg))
		return
	}

	err = job.conn.WriteMessage(websocket.BinaryMessage, flacData)
	if err != nil {
		log.Println("Write error:", err)
	}
}

func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

	cmd.Stdin = bytes.NewReader(wavData)
	var flacBuffer bytes.Buffer
	cmd.Stdout = &flacBuffer

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffmpeg error: %v", err)
	}

	return flacBuffer.Bytes(), nil
}
