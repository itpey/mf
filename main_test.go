package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	tests := []struct {
		filename    string
		expectError bool
	}{
		{"testfile1.txt", false},
		{"testfile1.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if _, err := os.Stat(tt.filename); err == nil {
				os.Remove(tt.filename)
			}
			file, err := os.OpenFile(tt.filename, os.O_CREATE|os.O_EXCL, 0666)
			if err != nil && !os.IsExist(err) && tt.expectError == false {
				t.Errorf("Error creating file: %v", err)
			}

			if err == nil && tt.expectError == false {
				defer file.Close()
				if _, err := os.Stat(tt.filename); os.IsNotExist(err) {
					t.Errorf("File '%s' was not created", tt.filename)
				}
			}

			if err != nil && tt.expectError && !os.IsExist(err) {
				t.Errorf("Expected error for existing file, but got: %v", err)
			}
		})
	}
}
