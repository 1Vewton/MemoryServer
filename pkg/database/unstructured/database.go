package unstructured

import (
	"context"

	"github.com/1Vewton/MemoryServer/pkg/memory"
)

// MemoryDB defines the interface of database for storing unstructured message memory
type MemoryDB interface {
	New(
		ctx context.Context,
		databaseURL string,
	) error
	NewMessage(
		ctx context.Context,
		msg *memory.MessageMemory,
	) error
}

// KnowledgeDB defines the interface of database for storing knowledge
type KnowledgeDB interface {
	New(
		ctx context.Context,
		databaseURL string,
	) error
	NewKnowledge(
		ctx context.Context,
		knowledge *memory.Knowledge,
	) error
}
