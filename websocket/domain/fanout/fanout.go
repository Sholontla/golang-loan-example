package fanout

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"github.com/projects/loans/utils/date_utils"
)

const (
	WorkersLenght = 4
)

func OperationA(a float64) float64 {
	return a * 10 / 2
}

func OperationB(b float64) float64 {
	return b * 10 / 2
}

func OperationC(c float64) float64 {
	return c * 10 / 2
}

func OperationD(d float64) float64 {
	return d * 10 / 2
}

func GetProcessor(a Worker1, b Worker2, c Worker3, d Worker4) *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers1:   make([]*Worker1, WorkersLenght),
		workers2:   make([]*Worker2, WorkersLenght),
		workers3:   make([]*Worker3, WorkersLenght),
		workers4:   make([]*Worker4, WorkersLenght),
		done1:      make(chan *Worker1),
		done2:      make(chan *Worker2),
		done3:      make(chan *Worker3),
		done4:      make(chan *Worker4),
	}
	for i := 0; i < WorkersLenght; i++ {
		a := &Worker1{ID: uuid.New(), Name: a.Name, Description: a.Description, FloatNum: OperationA(a.FloatNum), CreatedAt: date_utils.GetNowString()}
		b := &Worker2{ID: uuid.New(), Name: b.Name, Description: b.Description, FloatNum: OperationA(b.FloatNum), CreatedAt: date_utils.GetNowString()}
		c := &Worker3{ID: uuid.New(), Name: c.Name, Description: c.Description, FloatNum: OperationA(c.FloatNum), CreatedAt: date_utils.GetNowString()}
		d := &Worker4{ID: uuid.New(), Name: d.Name, Description: d.Description, FloatNum: OperationA(d.FloatNum), CreatedAt: date_utils.GetNowString()}

		filea, _ := json.MarshalIndent(a, "", " ")
		os.WriteFile("a.json", filea, 0777) //0644

		fileb, _ := json.MarshalIndent(b, "", " ")
		os.WriteFile("b.json", fileb, 0777) //0644

		filec, _ := json.MarshalIndent(c, "", " ")
		os.WriteFile("c.json", filec, 0777) //0644

		filed, _ := json.MarshalIndent(d, "", " ")
		os.WriteFile("d.json", filed, 0777) //0644

		p.workers1[i] = a
		p.workers2[i] = b
		p.workers3[i] = c
		p.workers4[i] = d
	}

	p.startProcess()

	return p
}

func (p *Processor) startProcess() {
	go func() {
		for {
			select {
			default:
				if len(p.workers1) > 0 {
					a := p.workers1[0]
					b := p.workers2[1]
					c := p.workers3[2]
					d := p.workers4[3]

					a.processJob1(<-p.jobChannel, p.done1)
					b.processJob2(<-p.jobChannel, p.done2)
					c.processJob3(<-p.jobChannel, p.done3)
					d.processJob4(<-p.jobChannel, p.done4)
				}
			case a := <-p.done1:
				p.workers1 = append(p.workers1, a)

			case b := <-p.done2:
				p.workers2 = append(p.workers2, b)

			case c := <-p.done3:
				p.workers3 = append(p.workers3, c)

			case d := <-p.done4:
				p.workers4 = append(p.workers4, d)
			}
		}
	}()

}
