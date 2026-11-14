package unstructured

import (
	"context"

	"github.com/1Vewton/MemoryServer/pkg/knowledge"
	"github.com/1Vewton/MemoryServer/pkg/memory"
)

// UnstructuredDB defines the connection for unstructured db
type UnstructuredDB interface {
	New(
		ctx context.Context,
		databaseURL string,
	) error
	Close(
		ctx context.Context,
	) error
}

// MemoryDB defines the interface of database for storing unstructured message memory
type MemoryDB interface {
	NewMessage(
		ctx context.Context,
		msg *memory.MessageMemory,
	) error
}

// KnowledgeDB defines the interface of database for storing knowledge
type KnowledgeDB interface {
	NewKnowledge(
		ctx context.Context,
		knowledge *knowledge.Piece,
	) error
}
