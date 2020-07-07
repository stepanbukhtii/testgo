package main

import (
	"fmt"

	"testgo"
)

type Pointer struct {
	Output1 *testgo.HardWork
	Output2 *testgo.HardWork
	Output3 *testgo.HardWork
	Output4 *testgo.HardWork
	Output5 *testgo.HardWork
	Counter int64
}

func (p *Pointer) Calc() {
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
	if p.Output5.GetValue() > p.Output5.GetValue() {
		p.Counter++
	}
}

func main() {
	p := testgo.MemTimeProfiler{}

	outputs1 := make([]testgo.HardWork, 40)
	outputs2 := make([]testgo.HardWork, 40)
	outputs3 := make([]testgo.HardWork, 40)
	outputs4 := make([]testgo.HardWork, 40)
	outputs5 := make([]testgo.HardWork, 40)

	p.Start()
	numberSubjects := make([]Pointer, 0, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4)*len(outputs5))
	for o1 := range outputs1 {
		for o2 := range outputs2 {
			for o3 := range outputs3 {
				for o4 := range outputs4 {
					for o5 := range outputs5 {
						numberSubjects = append(numberSubjects, Pointer{
							Output1: &outputs1[o1],
							Output2: &outputs2[o2],
							Output3: &outputs3[o3],
							Output4: &outputs4[o4],
							Output5: &outputs4[o5],
						})
					}
				}
			}
		}
	}
	p.Finish()

	fmt.Println("Size usage for init data", p.Size())

}
