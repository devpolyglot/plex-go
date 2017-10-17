package model

type Part struct {
	AudioProfile          string `json:"audioProfile"`
	Container             string `json:"container"`
	Duration              int    `json:"duration"`
	File                  string `json:"file"`
	Has64bitOffsets       bool   `json:"has64bitOffsets"`
	ID                    int    `json:"id"`
	Key                   string `json:"key"`
	OptimizedForStreaming bool   `json:"optimizedForStreaming"`
	Size                  int    `json:"size"`
	VideoProfile          string `json:"videoProfile"`
}
