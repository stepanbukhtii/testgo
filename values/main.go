package main

import (
	"fmt"
	"math/rand"

	"testgo"
)

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
			Output1: testgo.NewHardWork(i),
			Output2: testgo.NewHardWork(i),
			Output3: testgo.NewHardWork(i),
			Output4: testgo.NewHardWork(i),
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
			numberSubjects[y].Output1.Update(randomValues[i])
			numberSubjects[y].Output2.Update(randomValues[i])
			numberSubjects[y].Output3.Update(randomValues[i])
			numberSubjects[y].Output4.Update(randomValues[i])
			numberSubjects[y].Handle()
		}
	}
	p.Finish()
	fmt.Println("first loop valuer -> loop subjects", p.Time())

	p.Start()
	for _, subject := range numberSubjects {
		for i := 0; i < testgo.CountLoop; i++ {
			subject.Output1.Update(randomValues[i])
			subject.Output2.Update(randomValues[i])
			subject.Output3.Update(randomValues[i])
			subject.Output4.Update(randomValues[i])
			subject.Handle()
		}
	}
	p.Finish()

	fmt.Println("first loop subject -> loop values", p.Time())
}
