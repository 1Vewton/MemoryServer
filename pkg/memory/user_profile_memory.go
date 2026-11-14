package memory

// UserProfileMemory defines the memory for certain user
type UserProfileMemory struct {
	Nicknames  []string
	Hobbies    []string
	OtherInfos []string
}

// NewUserProfileMemory creates new user profile memory
func NewUserProfileMemory() *UserProfileMemory {
	return &UserProfileMemory{
		Nicknames:  []string{},
		Hobbies:    []string{},
		OtherInfos: []string{},
	}
}
