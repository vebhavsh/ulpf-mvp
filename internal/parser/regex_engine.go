package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"ulpf-go-engine/internal/models"
)

// Pre-compile regular expressions for maximum performance.
// In a real deployment, these patterns would be loaded from the PostgreSQL registry,
// but for the MVP, we are hardcoding the most common firewall patterns.
var (
	// Matches "src=192.168.1.1" or "srcip=10.0.0.1"
	reSrcIP = regexp.MustCompile(`(?:srcip|src)=([0-9\.]+)`)
	
	// Matches "dst=8.8.8.8" or "dstip=1.1.1.1"
	reDstIP = regexp.MustCompile(`(?:dstip|dst)=([0-9\.]+)`)
	
	// Matches "dstport=443" or "dpt=80" or "port=22"
	reDstPort = regexp.MustCompile(`(?:dstport|dpt|port)=([0-9]+)`)
	
	// Matches "action=allow" or "action=DENY"
	reAction = regexp.MustCompile(`action=([a-zA-Z]+)`)
)

// ParseNetworkLog takes a raw text log, applies regex patterns to extract fields,
// and returns a populated OCSF NetworkActivity struct.
func ParseNetworkLog(rawLog string) models.NetworkActivity {
	// Initialize the base OCSF structure with default metadata
	ocsfEvent := models.NetworkActivity{
		ClassUID:  4001,
		ClassName: "Network Activity",
		Time:      time.Now().UTC(), // Simulating log arrival time
		RawLog:    rawLog,           // Preserving the original log
	}

	// Extract Source IP using FindStringSubmatch
	// match[0] is the full match (e.g., "src=10.0.0.1")
	// match[1] is the captured group (e.g., "10.0.0.1")
	if match := reSrcIP.FindStringSubmatch(rawLog); len(match) > 1 {
		ocsfEvent.SrcEndpoint.IP = match[1]
	}

	// Extract Destination IP
	if match := reDstIP.FindStringSubmatch(rawLog); len(match) > 1 {
		ocsfEvent.DstEndpoint.IP = match[1]
	}

	// Extract Destination Port
	if match := reDstPort.FindStringSubmatch(rawLog); len(match) > 1 {
		// Convert string port to integer
		port, err := strconv.Atoi(match[1])
		if err == nil {
			ocsfEvent.DstEndpoint.Port = port
		}
	}

	// Extract Action
	if match := reAction.FindStringSubmatch(rawLog); len(match) > 1 {
		// Convert action to lowercase for standard normalization (e.g., "DENY" -> "deny")
		ocsfEvent.Action = strings.ToLower(match[1])
	}

	return ocsfEvent
}