package process

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenerateRandomNumbers(max int) (randomList []int) {

	randomList = rand.Perm(max)

	return
}