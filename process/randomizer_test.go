package process

import (
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	list1 := GenerateRandomNumbers(50)
	list2 := GenerateRandomNumbers(50)

	same := true
	for i := range list1 {
		if list1[i] != list2[i] {
			same = false
		}
	}
	if same {
		t.Error("Generated 2 lists with same number pattern")
	}
}
