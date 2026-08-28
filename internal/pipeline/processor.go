package pipeline

import (
	"fmt"
	"ulpf-go-engine/internal/database"

	"ulpf-go-engine/internal/normalizer"
	"ulpf-go-engine/internal/parser"
)

// RunPipeline is the core function that coordinates the fast data plane.
// It takes the raw log string and passes it through the parsing and normalization stages.
func RunPipeline(rawLog string, sourceIP string) (string, error) {
	// Step 1: Pass the raw text to the Regex Engine to extract fields
	parsedEvent := parser.ParseNetworkLog(rawLog)

	// Step 2: Pass the struct to the Normalizer for validation and JSON conversion
	finalJSONBytes, err := normalizer.FinalizeEvent(parsedEvent)
	if err != nil {
		// If validation fails
		return "", fmt.Errorf("pipeline rejected log from %s: %v", sourceIP, err)
	}

	// Bytes ko string mein convert kar rahe hain taaki database mein save ho sake
	finalJSONString := string(finalJSONBytes)

	// Step 3: Save to Database! (Ye tera naya addition hai)
	err = database.InsertLog(rawLog, finalJSONString)
	if err != nil {
		fmt.Printf("Warning: Failed to save to database: %v\n", err)
	}

	// Step 4: Return the final, successfully normalized JSON string
	return finalJSONString, nil
}
