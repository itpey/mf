package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		return
	}

	fileName := os.Args[1]

	if err := validateFilename(fileName); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("File '%s' already exists.\n", fileName)
		} else {
			fmt.Printf("Error creating file: %v\n", err)
		}
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' created successfully.\n", fileName)
}

func validateFilename(fileName string) error {

	if fileName == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	if runtime.GOOS == "windows" {

		invalidChars := `[<>:"/\\|?*]`
		re := regexp.MustCompile(invalidChars)
		if re.MatchString(fileName) {
			return fmt.Errorf("filename contains invalid characters for Windows")
		}

		if strings.HasSuffix(fileName, " ") || strings.HasSuffix(fileName, ".") {
			return fmt.Errorf("filename cannot end with a space or period on Windows")
		}
	} else {
		invalidChars := `[\x00/]`
		re := regexp.MustCompile(invalidChars)
		if re.MatchString(fileName) {
			return fmt.Errorf("filename contains invalid characters (null byte or /) for Unix-based systems")
		}
	}

	if len(fileName) > 255 {
		return fmt.Errorf("filename is too long (max 255 characters)")
	}

	if strings.TrimSpace(fileName) == "" {
		return fmt.Errorf("filename cannot be just spaces")
	}

	return nil
}
