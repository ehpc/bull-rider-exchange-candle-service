package testing

import (
	"os"
	"testing"
	"log"

	"github.com/joho/godotenv"
)

// Setup initialized testing environment
func Setup(m *testing.M) {
	// Load configuration
	err := godotenv.Load("../../.env")
	if err != nil {
		dir, _ := os.Getwd()
		log.Fatal(err, dir)
	}
	os.Exit(m.Run())
}
