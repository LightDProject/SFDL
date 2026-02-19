package sfdl

import (
	"testing"
)

func TestParse(t *testing.T) {
	config := Config{
		Filename: "test.sfdl",
		Content:  []byte(`name = "test"`),
	}

	file, err := Parse(config)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if file == nil {
		t.Fatal("File is nil")
	}
}
