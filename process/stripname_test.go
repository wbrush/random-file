package process

import (
	"testing"
	"strings"
)

func TestStripNames(t *testing.T) {
	var testIn []string

	if testIn == nil {
		testIn = make([]string, 0)
	}

	testIn = append(testIn, "1_file1")
	testIn = append(testIn, "2098_file2")
	testIn = append(testIn, "file3")

	testOut, err := StripNames(testIn)
	if err != nil {
		t.Error("Received error in StripNames:", err.Error())
	}
	if len(testOut) != 3 {
		t.Error("Invalid number of files returned")
	}
	if strings.Compare(testOut[0], "file1") != 0 {
		t.Error("File 1 doesn't match expected name")
	}
	if strings.Compare(testOut[1], "file2") != 0 {
		t.Error("File 2 doesn't match expected name")
	}
	if strings.Compare(testOut[2], "file3") != 0 {
		t.Error("File 3 doesn't match expected name")
	}
}