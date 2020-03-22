package main

import "testing"

func TestIsValidToDecompress_Sucess(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
	}{
		{"should pass on validation with regular name", "file.huff"},
		{"should pass on validation with one char file name", "a.huff"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidToDecompress(tt.fileName)
			if err != nil {
				t.Errorf("got %q, want nil", err)
			}
		})
	}
}

func TestIsValidToDecompress_Error(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
		out      string
	}{
		{"should fail due to file name has no .huff", "payments.csv", "expected extension huff, given csv"},
		{"should fail due to file name has no extension", "filename", "file without extension: filename"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidToDecompress(tt.fileName)
			if err.Error() != tt.out {
				t.Errorf("got %q, want %s", err, tt.out)
			}
		})
	}
}
