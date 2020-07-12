package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/stepanbukhtii/testgo"
)

// Size usage for init data 800.0 MB
// Using value loop time usage 42.5570102s

type ArgumentsSingle struct {
	Counter int64
}

func (a *ArgumentsSingle) Handle(value testgo.TestStructure) {
	if value.Test1 > a.Counter {
		a.Counter++
	}
	if value.Test2 > a.Counter {
		a.Counter++
	}
	if value.Test3 > a.Counter {
		a.Counter++
	}
	if value.Test4 > a.Counter {
		a.Counter++
	}
}

func main() {
	p := testgo.MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	hw1 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw2 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw3 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw4 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))

	numberSubjects := make([]ArgumentsSingle, 0, len(hw1)*len(hw2)*len(hw3)*len(hw4))
	p.Finish()

	fmt.Println("Size usage for init data", p.Size())

	p.Start()
	for _, v1 := range hw1 {
		for _, v2 := range hw2 {
			for _, v3 := range hw3 {
				for _, v4 := range hw4 {
					for i := 0; i < testgo.CountLoop; i++ {
						value := testgo.TestStructure{
							Test1: v1[i],
							Test2: v2[i],
							Test3: v3[i],
							Test4: v4[i],
						}
						for y := range numberSubjects {
							numberSubjects[y].Handle(value)
						}
					}
				}
			}
		}
	}
	p.Finish()

	fmt.Println("Using value loop time usage", p.Time())
}
