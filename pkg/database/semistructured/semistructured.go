package semistructured

import (
	"context"

	"github.com/1Vewton/MemoryServer/pkg/memory"
)

// UserProfileDB defines the interface for storing user profile
type UserProfileDB interface {
	New(
		ctx context.Context,
		databaseURL string,
	) error
	NewUserProfile(
		ctx context.Context,
		userProfile *memory.UserProfileMemory,
	) error
}
