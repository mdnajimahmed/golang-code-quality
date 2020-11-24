package utility

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
)

func MakeRandInt(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100000)
	}
	return numbers
}

func ReadHttpBody(body io.ReadCloser) string {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}
