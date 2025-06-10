package utils

import (
	"encoding/json"
	"log"
)

func MustMarshal(v any) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}
	return data
}

func MustUnmarshal(data []byte, v any) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}
}
