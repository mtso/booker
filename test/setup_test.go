package test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// drop tables and create test users 1 and 2
	code := m.Run()
	// drop tables
	os.Exit(code)
}
