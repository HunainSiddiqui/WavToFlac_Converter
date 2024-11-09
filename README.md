# Audio Conversion Service

This project provides a real-time audio conversion service that converts WAV audio streams to FLAC format using Go, Gorilla Mux, and WebSockets. The service can handle multiple simultaneous audio streams efficiently.

## LIVE DEMO

The client application is hosted on Vercel and can be accessed at the following URL:

[Live Demo](https://front-task-zeta.vercel.app/)

This is just a demo we deployed the backend on render here is the link to the backend

https://wavtoflac-converter.onrender.com

this is the link to the backend you can test the backend by sending the audio file to the backend and it will convert the audio file to flac format and send it back to you.


## Features

- Real-time audio streaming and conversion
- WebSocket support for seamless communication
- CORS middleware for cross-origin requests
- Scalable architecture with worker pool for concurrent processing

## Technologies Used

- Go
- Gorilla Mux
- WebSocket
- FFmpeg (for audio conversion)


## Getting Started

### Prerequisites

- Go (version 1.18 or later)
- FFmpeg (make sure it's installed and accessible in your PATH)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/HunainSiddiqui/Peritys_TAsk.git

   cd Peritys_TAsk
    ```
2. Run the application:

   ```bash
   go mod tidy
   go run main.go
   ```
3. The server will start on port 300 by default. 

4. Open the client application in your browser:

   ```
   http//localhost:3000
   ```
5. Start streaming audio and see the conversion in real-time!

## Testing

### Unit Tests

Run the unit tests to validate the conversion logic:

```bash
go test ./tests.
```

## Integration Tests

You can run integration tests to check WebSocket connections and the real-time streaming capability.We have use testify for the integration test.

```bash
    go test ./tests 
```

## API Endpoints

### WebSocket Endpoint

- **Endpoint:** `/ws/{id}`
- **Method:** `GET`
- **Description:** Establishes a WebSocket connection to convert audio streams. The `{id}` parameter represents a unique identifier for the audio stream.
  
- **Request:**
  - Send WAV audio data as a binary message to the WebSocket connection.

- **Response:**
  - The service returns the converted FLAC audio data as a binary message.


