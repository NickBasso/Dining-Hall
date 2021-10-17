package configs

import "os"

func SetupENV() {
	os.Setenv("DHALL_URL", "http://localhost:4005")
	os.Setenv("KITCHEN_URL", "http://localhost:4006")
}