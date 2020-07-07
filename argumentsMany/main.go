package main

import (
	"fmt"
	"math/rand"
	"time"

	"testgo"
)

// Size usage for init data 358.4 kB
// Using value loop time usage 2m56.297047654s

func main() {
	p := testgo.MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	hw1 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw2 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw3 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw4 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	p.Finish()

	fmt.Println("Size usage for init data", p.Size())

	var counter int64
	var first, second, third testgo.TestStructure

	p.Start()
	for _, v1 := range hw1 {
		for _, v2 := range hw2 {
			for _, v3 := range hw3 {
				for _, v4 := range hw4 {
					for i := 0; i < testgo.CountLoop; i++ {
						first = second
						second = third
						third = testgo.TestStructure{
							Test1: v1[i],
							Test2: v2[i],
							Test3: v3[i],
							Test4: v4[i],
						}
						r := HandleWork(first, second, third)
						counter = counter + r
					}
				}
			}
		}
	}
	p.Finish()
	fmt.Println("Using value loop time usage", p.Time())

	fmt.Println("counter", counter)
}

func HandleWork(v1, v2, v3 testgo.TestStructure) int64 {
	counter := int64(0)
	if v1.Test1 > v2.Test1 {
		counter++
	}
	if v1.Test2 > v2.Test2 {
		counter++
	}
	if v1.Test3 > v2.Test3 {
		counter++
	}
	if v1.Test4 > v3.Test4 {
		counter++
	}
	return counter
}
