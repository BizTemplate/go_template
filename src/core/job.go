package core

import "fmt"

type Job interface {
	Perform()
}
type SimpleJob struct {
	name string
}

func (job *SimpleJob) Perform() {
	fmt.Printf("start")
	job.initData()
}

func (job *SimpleJob) initData() {
	fmt.Printf("start")
}
