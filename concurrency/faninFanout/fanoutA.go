package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func GenerateA(data string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- data
			time.Sleep(1000)
		}

	}()
	return channel
}

type ProcessorA struct {
	jobChannel chan string
	done       chan *WorkerA
	done2      chan *Worker2A
	workers    []*Worker2A
	workers1   []*WorkerA
	workers2   []*WorkerA
	workers3   []*WorkerA
}
type WorkerA struct {
	Name     string  `json:"name"`
	Name2    string  `json:"name2"`
	FloatNum float64 `json:"float_num"`
}

type Worker2A struct {
	Name     string  `json:"name"`
	Name2    string  `json:"name2"`
	FloatNum float64 `json:"float_num"`
}

func (w *WorkerA) processJobA(data string, done chan *WorkerA) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, w.Name)
		fmt.Println("Working on data 2", data, w.Name2)
		fmt.Println("Working on floatNum 1", data, w.FloatNum)
		fmt.Printf("\n")
		done <- w
	}()
}

func (x *Worker2A) processJob2A(data string, done chan *Worker2A) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, x.Name)
		fmt.Println("Working on data 2", data, x.Name2)
		fmt.Println("Working on floatNum 1", data, x.FloatNum)
		fmt.Printf("\n")
		done <- x
	}()
}

func GetProcessorA(a Worker2A) *ProcessorA {
	p := &ProcessorA{
		jobChannel: make(chan string),
		workers:    make([]*Worker2A, WorkersLenght),
		workers1:   make([]*WorkerA, WorkersLenght),
		workers2:   make([]*WorkerA, WorkersLenght),
		workers3:   make([]*WorkerA, WorkersLenght),
		done:       make(chan *WorkerA),
		done2:      make(chan *Worker2A),
	}

	for i := 0; i < WorkersLenght; i++ {

		w := &Worker2A{Name: a.Name, Name2: a.Name2, FloatNum: a.FloatNum * 10 / 2}
		filew, _ := json.MarshalIndent(w, "", " ")
		os.WriteFile("w.json", filew, 0777) //0644

		x := &WorkerA{Name: a.Name, Name2: a.Name, FloatNum: a.FloatNum * 10 / 2}
		filex, _ := json.MarshalIndent(x, "", " ")
		os.WriteFile("x.json", filex, 0777) //0644

		y := &WorkerA{Name: a.Name, Name2: a.Name, FloatNum: a.FloatNum * 10 / 2}
		filey, _ := json.MarshalIndent(y, "", " ")
		os.WriteFile("y.json", filey, 0777) //0644

		z := &WorkerA{Name: a.Name, Name2: a.Name, FloatNum: a.FloatNum * 10 / 2}
		filez, _ := json.MarshalIndent(z, "", " ")
		os.WriteFile("z.json", filez, 0777) //0644

		p.workers[i] = w
		p.workers1[i] = x
		p.workers2[i] = y
		p.workers3[i] = z
	}

	p.startProcessA()

	return p
}

func (p *ProcessorA) startProcessA() {
	go func() {
		for {
			select {
			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					x := p.workers1[1]
					y := p.workers2[2]
					z := p.workers3[3]
					p.workers = p.workers[0:]
					w.processJob2A(<-p.jobChannel, p.done2)
					x.processJobA(<-p.jobChannel, p.done)
					y.processJobA(<-p.jobChannel, p.done)
					z.processJobA(<-p.jobChannel, p.done)

				}
			case w := <-p.done2:
				p.workers = append(p.workers, w)

			case x := <-p.done:
				p.workers1 = append(p.workers1, x)

			case y := <-p.done:
				p.workers2 = append(p.workers2, y)

			case z := <-p.done:
				p.workers3 = append(p.workers3, z)
			}
		}
	}()
}

func (p *ProcessorA) PostJobA(jobs <-chan string) {
	p.jobChannel <- <-jobs
	wg.Done()

}
