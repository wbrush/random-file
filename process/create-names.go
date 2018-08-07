package process

import (
	"errors"
	"fmt"
)

func CreateNames (files []string, nums []int) (newList []string, err error) {

	//  validate length of both arrays\
	if (len(files) != len(nums)) {
		err = errors.New("Array Length Mismatch")
		return
	}

	//  for each file
	for i := range files {

		//  create new file name
		item := fmt.Sprintf("%03d_%s", nums[i], files[i])
		newList = append(newList, item)

	}

	return
}
