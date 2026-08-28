package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"ulpf-go-engine/internal/database"
	"ulpf-go-engine/internal/pipeline"
)

// LogEnvelope represents the initial wrapper for incoming raw logs.
// It stores the raw string and metadata before parsing begins.
type LogEnvelope struct {
	RawData string
	Source  string
}

// processLog simulates the fast data plane pipeline.
// This function is designed to run as an independent goroutine.
func processLog(envelope LogEnvelope, resultChan chan string) {
	// Ab hum dummy text ki jagah asli pipeline call kar rahe hain
	finalJSON, err := pipeline.RunPipeline(envelope.RawData, envelope.Source)

	if err != nil {
		// Agar validation fail hui (e.g. invalid log)
		resultChan <- fmt.Sprintf("DLQ ALERT: %v\nRaw Log: %s", err, envelope.RawData)
		return
	}

	// Agar sab pass ho gaya toh JSON return karega
	resultChan <- fmt.Sprintf("Successfully Normalized Event:\n%s", finalJSON)
}

// ingestHandler is the HTTP endpoint that receives incoming logs from Fluent Bit or direct POST requests.
func ingestHandler(w http.ResponseWriter, r *http.Request) {
	// Enforce POST method for log submission
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Read the raw log from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the request body.", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	rawLog := string(body)
	if rawLog == "" {
		http.Error(w, "Log body cannot be empty.", http.StatusBadRequest)
		return
	}

	// Create a channel to handle communication between the HTTP handler and the worker goroutine
	resultChan := make(chan string)

	// Wrap the raw log into an envelope with source metadata
	envelope := LogEnvelope{
		RawData: rawLog,
		Source:  r.RemoteAddr,
	}

	// Launch a goroutine to process the log asynchronously without blocking the server
	go processLog(envelope, resultChan)

	// Wait for the goroutine to finish processing and send the output back
	result := <-resultChan

	// Send the HTTP 200 OK response back to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", result)
}

func main() {
	database.InitDB()
	// Register the /ingest endpoint to our handler function
	http.HandleFunc("/ingest", ingestHandler)

	port := ":8080"
	fmt.Printf("ULPF Go Engine is starting on port %s...\n", port)

	// Start the HTTP server and log any fatal errors if it crashes
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
