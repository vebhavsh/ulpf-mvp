package normalizer

import (
	"encoding/json"
	"errors"
	"fmt"

	"ulpf-go-engine/internal/models"
)

// FinalizeEvent takes the parsed OCSF struct, validates it against core rules,
// and normalizes it into the final JSON payload ready for the Data Lake / SIEM.
func FinalizeEvent(event models.NetworkActivity) ([]byte, error) {
	// ---------------------------------------------------------
	// STEP 1: Schema Validation (Crucial for Blueprint)
	// ---------------------------------------------------------
	// Agar log itna malformed hai ki usme Source IP ya Destination IP dono hi nahi hain,
	// toh ye OCSF standard ke khilaf hai. Hum isko error ke sath return karenge
	// taaki isko "Dead-Letter Queue (DLQ)" mein bheja ja sake.
	if event.SrcEndpoint.IP == "" && event.DstEndpoint.IP == "" {
		return nil, errors.New("schema validation failed: both Source and Destination IPs are missing")
	}

	// ---------------------------------------------------------
	// STEP 2: Normalization to JSON
	// ---------------------------------------------------------
	// Hackathon ke demo ke liye hum 'MarshalIndent' use kar rahe hain taaki 
	// terminal pe JSON ekdum sundar aur readable format mein print ho.
	// (Production mein high-speed ke liye hum normal 'json.Marshal' use karte hain).
	finalJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("JSON serialization failed: %v", err)
	}

	return finalJSON, nil
}