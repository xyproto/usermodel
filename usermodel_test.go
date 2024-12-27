package usermodel

import (
	"fmt"
	"testing"
)

func TestGetCodeCompletionModel(t *testing.T) {
	ccm := GetCodeCompletionModel()
	if ccm == "" {
		fmt.Println("Warning: the code completion model is empty. Is llm-manager installed?")
	}
	fmt.Printf("code completion model = %s\n", GetCodeCompletionModel())
}
