package models

import "time"

// Endpoint represents a network endpoint (Source or Destination) in OCSF.
type Endpoint struct {
	IP   string `json:"ip"`
	Port int    `json:"port,omitempty"`
}

// NetworkActivity represents the OCSF Network Activity schema (Class UID 4001).
// This struct defines how our normalized JSON output will look.
type NetworkActivity struct {
	// Standard OCSF Metadata
	ClassUID  int       `json:"class_uid"`
	ClassName string    `json:"class_name"`
	Time      time.Time `json:"time"`

	// Extracted Network Data
	SrcEndpoint Endpoint `json:"src_endpoint"`
	DstEndpoint Endpoint `json:"dst_endpoint"`

	// Action taken by the firewall (e.g., "Allow", "Deny")
	Action string `json:"action,omitempty"`

	// Preserving the original raw log for traceability (Crucial ULPF Requirement)
	RawLog string `json:"raw_log"`
}