package settings

import (
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
}

// Initialize reads the env file setted
func (s *settings) Initialize(filePath string) error {
	err := godotenv.Load(filePath)
	return err
}
