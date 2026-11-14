package memory

import (
	"testing"
)

// TestMessageMemory tests functions related to message memory
func TestMessageMemory(t *testing.T) {
	t.Parallel()
	newMessage := NewMessage(
		"abc",
		"cde",
	).SwitchSharedState()
	if !newMessage.Shared {
		t.Fatal(
			"expected the message shared status to be true",
		)
	}
}
