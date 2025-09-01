package config

import (
	"os"
)

func Port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8085"
}
