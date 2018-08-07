package os_services

import (
	"testing"
)

func TestReadDirectory(t *testing.T) {

	dir, err := ReadDirectory(".")
	if err != nil {
		t.Error("Got error reading root directory:", err.Error())
	}
	if len(dir) != 4 {
		t.Error("Wrong number of files in root directory:", len(dir))
	}
}

func TestReadDirectory2(t *testing.T) {

	_, err := ReadDirectory("./unknown_directory")
	if err == nil {
		t.Error("Didn't get an error reading non-existent directory")
	}
}
