package main

import (
	"fmt"

	"testgo"
)

type ArgumentsSingle struct {
	Counter int64
}

func (a *ArgumentsSingle) ArgumentFuncSingle(value testgo.TestStructure) {
	if a.Counter > value.Test1 {
		a.Counter++
	}
	if a.Counter > value.Test2 {
		a.Counter++
	}
	if a.Counter > value.Test3 {
		a.Counter++
	}
	if a.Counter > value.Test4 {
		a.Counter++
	}
}

func main() {
	p := testgo.MemTimeProfiler{}

	randomValues := testgo.GenerateRandomTestStructures(testgo.TenMillions)

	p.Start()
	hw1 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw2 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw3 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	hw4 := testgo.FillSliceOfInt(make([][]int64, testgo.CountLoop))
	p.Finish()

	p.Start()
	numberSubjects := make([]ArgumentsSingle, 0, len(hw1)*len(hw2)*len(hw3)*len(hw4))
	for o1 := range hw1 {
		for o2 := range hw2 {
			for o3 := range hw3 {
				for o4 := range hw4 {
					numberSubjects = append(numberSubjects, ArgumentsSingle{})
				}
			}
		}
	}
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

}
