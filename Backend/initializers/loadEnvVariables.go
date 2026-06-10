package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file, this is not a problem if its currently running in a docker container unless the program stops.")
	}
}
