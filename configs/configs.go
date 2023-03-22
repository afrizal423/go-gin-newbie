package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// memanggil konfig pada env file
func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading file env")
	}

	return os.Getenv(key)
}
