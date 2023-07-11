package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	envs := []string{
		"mysql",
		"mariadb",
		//"sqlite3",
	}
	os.Setenv("DEBUG", "true")
	for _, env := range envs {
		t.Logf("------------ Start Testing %s ------------\n", env)
		os.Setenv("GORM_DIALECT", env)

		if env == "sqlite3" {
			os.Setenv("CGO_ENABLED", "0")
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("PANIC: %v\n", r)
				}
			}()
			main()
		}()

		// Reset the environment variable.
		os.Unsetenv("GORM_DIALECT")
		t.Logf("------------- End Testing %s -------------\n", env)
	}
}
