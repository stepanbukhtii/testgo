package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/stepanbukhtii/testgo"
)

// Size usage for init data 4.0 GB
// first loop value -> loop subjects 55.0379986s

type Pointer struct {
	Output1 *testgo.HardWork
	Output2 *testgo.HardWork
	Output3 *testgo.HardWork
	Output4 *testgo.HardWork
	Counter int64
}

func (p *Pointer) Handle() {
	if p.Output1.GetValue() > p.Output4.GetValue() {
		p.Counter++
	}
	if p.Output2.GetValue() > p.Output1.GetValue() {
		p.Counter++
	}
	if p.Output3.GetValue() > p.Output2.GetValue() {
		p.Counter++
	}
	if p.Output4.GetValue() > p.Output3.GetValue() {
		p.Counter++
	}
}

func main() {
	p := testgo.MemTimeProfiler{}

	outputs1 := make([]testgo.HardWork, 100)
	for i := range outputs1 {
		outputs1[i] = testgo.NewHardWork(i + 1)
	}
	outputs2 := make([]testgo.HardWork, 100)
	for i := range outputs2 {
		outputs2[i] = testgo.NewHardWork(i + 1)
	}
	outputs3 := make([]testgo.HardWork, 100)
	for i := range outputs3 {
		outputs3[i] = testgo.NewHardWork(i + 1)
	}
	outputs4 := make([]testgo.HardWork, 100)
	for i := range outputs4 {
		outputs4[i] = testgo.NewHardWork(i + 1)
	}

	p.Start()
	numberSubjects := make([]Pointer, 0, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	for o1 := range outputs1 {
		for o2 := range outputs2 {
			for o3 := range outputs3 {
				for o4 := range outputs4 {
					numberSubjects = append(numberSubjects, Pointer{
						Output1: &outputs1[o1],
						Output2: &outputs2[o2],
						Output3: &outputs3[o3],
						Output4: &outputs4[o4],
					})
				}
			}
		}
	}
	p.Finish()

	fmt.Println("Size usage for init data", p.Size())

	rand.Seed(time.Now().UnixNano())

	randomValues := make([]int64, testgo.CountLoop)
	for i := range randomValues {
		randomValues[i] = rand.Int63n(10000)
	}

	p.Start()
	for i := 0; i < testgo.CountLoop; i++ {
		value := randomValues[i]
		for o := range outputs1 {
			outputs1[o].Update(value)
		}
		for o := range outputs2 {
			outputs2[o].Update(value)
		}
		for o := range outputs3 {
			outputs3[o].Update(value)
		}
		for o := range outputs4 {
			outputs4[o].Update(value)
		}

		for y := range numberSubjects {
			numberSubjects[y].Handle()
		}
	}
	p.Finish()
	fmt.Println("first loop value -> loop subjects", p.Time())
}
