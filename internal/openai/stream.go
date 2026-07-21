package openai

import "encoding/json"

type streamEvent struct {
	Type  string `json:"type"`
	Delta string `json:"delta,omitempty"`
}

func parseStreamEvent(data []byte) (streamEvent, error) {
	var event streamEvent

	if err := json.Unmarshal(data, &event); err != nil {
		return streamEvent{}, err
	}

	return event, nil
}
