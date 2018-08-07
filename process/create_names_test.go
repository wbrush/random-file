package process

import (
	"testing"
	"strings"
)

func getTestData() (files []string, nums []int) {
	if files == nil { files = make([]string, 0) }
	if nums  == nil { nums  = make([]int, 0) }

	files = append(files, "file1")
	files = append(files, "file2")
	files = append(files, "file3")

	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 1)

	return
}

func TestCreateNames(t *testing.T) {
	//  create input data lists - happy path
	files, nums := getTestData()

	//  call routine
	testOut, err := CreateNames(files, nums)
	if err != nil {
		t.Error("Got error calling CreateNames():", err.Error())
	}
	if len(testOut) != 3 {
		t.Error("Invalid number of items returned:", len(testOut))
	}

	//  verify that routine renames the files appropriately
	if strings.Compare(testOut[0], "003_file1") != 0 {
		t.Error("File name 0 doesn't match")
	}
	if strings.Compare(testOut[1], "002_file2") != 0 {
		t.Error("File name 1 doesn't match")
	}
	if strings.Compare(testOut[2], "001_file3") != 0 {
		t.Error("File name 2 doesn't match")
	}
}

func TestCreateNames2(t *testing.T) {
	//  create input data lists - mismatch in length of arrays (nums shorter)
	files, nums := getTestData()
	files = append(files, "file4")

	//  call routine
	_, err := CreateNames(files, nums)

	//  verify that routine returns an error
	if err == nil {
		t.Error("Didn't get an error with mismatched data lengths")
	}
}
