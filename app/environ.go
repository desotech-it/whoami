package app

import (
	"os"
	"sort"
)

// GetEnvironment returns all environment variables sorted alphabetically.
func GetEnvironment() []string {
	env := os.Environ()
	sort.Strings(env)
	return env
}
