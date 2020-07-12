package main

import (
	"fmt"
	"math/rand"

	"github.com/stepanbukhtii/testgo"
)

// Size usage for init data 10.4 GB
// first loop valuer -> loop subjects 5m25.1937319s
// first loop subject -> loop values 5m21.5470585s

type Value struct {
	Output1 testgo.HardWork
	Output2 testgo.HardWork
	Output3 testgo.HardWork
	Output4 testgo.HardWork
	Counter int64
}

func (p *Value) Handle() {
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

	p.Start()
	numberSubjects := make([]Value, 100*100*100*100)
	for i := range numberSubjects {
		numberSubjects[i] = Value{
			Output1: testgo.NewHardWork(i % 10 + 10),
			Output2: testgo.NewHardWork(i % 10 + 10),
			Output3: testgo.NewHardWork(i % 10 + 10),
			Output4: testgo.NewHardWork(i % 10 + 10),
		}
	}
	p.Finish()
	fmt.Println("Size usage for init data", p.Size())

	randomValues := make([]int64, testgo.CountLoop)
	for i := range randomValues {
		randomValues[i] = rand.Int63n(10000)
	}

	p.Start()
	for i := 0; i < testgo.CountLoop; i++ {
		for y := range numberSubjects {
			value := randomValues[i]
			numberSubjects[y].Output1.Update(value)
			numberSubjects[y].Output2.Update(value)
			numberSubjects[y].Output3.Update(value)
			numberSubjects[y].Output4.Update(value)
			numberSubjects[y].Handle()
		}
	}
	p.Finish()
	fmt.Println("first loop valuer -> loop subjects", p.Time())

	p.Start()
	for _, subject := range numberSubjects {
		for i := 0; i < testgo.CountLoop; i++ {
			value := randomValues[i]
			subject.Output1.Update(value)
			subject.Output2.Update(value)
			subject.Output3.Update(value)
			subject.Output4.Update(value)
			subject.Handle()
		}
	}
	p.Finish()

	fmt.Println("first loop subject -> loop values", p.Time())
}
