package testgo

import (
	"math/rand"
	"time"
)

const (
	Million     = 1000000
	TenMillions = 10 * Million
	CountLoop   = 100
)

type TestStructure struct {
	Test1  int64
	Test2  int64
	Test3  int64
	Test4  int64
	Test5  int64
	Test6  int64
	Test7  int64
	Test8  int64
	Test9  int64
	Test10 int64
}

func GenerateRandomTestStructures(quantity int) []TestStructure {
	rand.Seed(time.Now().UnixNano())

	randomValues := make([]TestStructure, quantity)
	for y := range randomValues {
		randomValues[y] = TestStructure{
			Test1:  rand.Int63n(10000),
			Test2:  rand.Int63n(10000),
			Test3:  rand.Int63n(10000),
			Test4:  rand.Int63n(10000),
			Test5:  rand.Int63n(10000),
			Test6:  rand.Int63n(10000),
			Test7:  rand.Int63n(10000),
			Test8:  rand.Int63n(10000),
			Test9:  rand.Int63n(10000),
			Test10: rand.Int63n(10000),
		}
	}

	return randomValues
}

func FillSliceOfInt(outputs [][]int64) [][]int64 {
	for i := range outputs {
		randomValues := make([]int64, CountLoop)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs[i] = randomValues
	}
	return outputs
}
