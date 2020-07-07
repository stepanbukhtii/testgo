package main

import (
	"fmt"
	"math/rand"
	"time"

	"testgo"
)

// Size usage 358.4 kB 0 B time usage 1m40.233014837s
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
	argumentValue := make([]testgo.TestStructure, 3)

	p.Start()
	for _, v1 := range hw1 {
		for _, v2 := range hw2 {
			for _, v3 := range hw3 {
				for _, v4 := range hw4 {
					for i := 0; i < testgo.CountLoop; i++ {
						argumentValue[2] = testgo.TestStructure{
							Test1: v1[i],
							Test2: v2[i],
							Test3: v3[i],
							Test4: v4[i],
						}
						r := HandleWork(argumentValue)
						argumentValue[0] = argumentValue[1]
						argumentValue[1] = argumentValue[2]
						counter = counter + r
					}
				}
			}
		}
	}
	p.Finish()
	fmt.Println("Using value loop time usage", p.Time())

	p.Start()
	for i1 := range hw1 {
		for i2 := range hw2 {
			for i3 := range hw3 {
				for i4 := range hw4 {
					for i := 0; i < testgo.CountLoop; i++ {
						argumentValue[2] = testgo.TestStructure{
							Test1: hw1[i1][i],
							Test2: hw2[i2][i],
							Test3: hw3[i3][i],
							Test4: hw4[i4][i],
						}
						r := HandleWork(argumentValue)
						argumentValue[0] = argumentValue[1]
						argumentValue[1] = argumentValue[2]
						counter = counter + r
					}
				}
			}
		}
	}
	p.Finish()
	fmt.Println("Using index loop time usage", p.Time())

	p.Start()

	for i := 0; i < testgo.CountLoop; i++ {
		for i1 := range hw1 {
			for i2 := range hw2 {
				for i3 := range hw3 {
					for i4 := range hw4 {
						argumentValue[2] = testgo.TestStructure{
							Test1: hw1[i1][i],
							Test2: hw2[i2][i],
							Test3: hw3[i3][i],
							Test4: hw4[i4][i],
						}
						r := HandleWork(argumentValue)
						argumentValue[0] = argumentValue[1]
						argumentValue[1] = argumentValue[2]
						counter = counter + r
					}
				}
			}
		}
	}
	p.Finish()
	fmt.Println("Using index loop first count loop time usage", p.Time())
}

func HandleWork(values []testgo.TestStructure) int64 {
	counter := int64(0)
	if values[0].Test1 > values[1].Test1 {
		counter++
	}
	if values[0].Test2 > values[1].Test2 {
		counter++
	}
	if values[0].Test3 > values[1].Test3 {
		counter++
	}
	if values[0].Test4 > values[1].Test4 {
		counter++
	}
	return counter
}
