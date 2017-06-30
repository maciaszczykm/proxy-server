package main

import (
	"strings"
	"encoding/json"
	"log"
	"fmt"
)

const (
	whitespace     = " "
	ackMessageType = "ACK"
	reqMessageType = "REQ"
	nakMessageType = "NAK"
)

// Stats of proxy traffic.
type Stats struct {
	TotalMessages int `json:"msg_total"`
	ReqMessages   int `json:"msg_req"`
	AckMessages   int `json:"msg_ack"`
	NakMessages   int `json:"msg_nak"`
}

// NewStats constructs new Stats object.
func NewStats() *Stats {
	return &Stats{
		TotalMessages: 0,
		ReqMessages:   0,
		AckMessages:   0,
		NakMessages:   0,
	}
}

// Update statistics processing single message passed through proxy.
func (s *Stats) Update(message string) {
	parts := strings.Split(message, whitespace)
	if len(parts) > 1 {
		s.TotalMessages += 1
		messageType := parts[0]
		switch messageType {
		case ackMessageType:
			s.AckMessages += 1
		case reqMessageType:
			s.ReqMessages += 1
		case nakMessageType:
			s.NakMessages += 1
		}
	}
}

// Print current proxy traffic statistics in JSON format.
func (s *Stats) Print() {
	bytes, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bytes))
}
