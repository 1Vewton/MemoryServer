package memory

import (
	"time"
)

// MessageMemory defines the memory for storing messages
type MessageMemory struct {
	Sender  string
	Content string
	Time    time.Time
	Shared  bool
}

// NewMessage creates new message memory
func NewMessage(
	sender string,
	content string,
) *MessageMemory {
	return &MessageMemory{
		Sender:  sender,
		Content: content,
		Time:    time.Now(),
		Shared:  false,
	}
}

// SwithcSharedState switches the shared state
func (msg *MessageMemory) SwitchSharedState() *MessageMemory {
	msg.Shared = !msg.Shared
	return msg
}
