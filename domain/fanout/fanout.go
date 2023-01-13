package fanout

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"github.com/projects/loans/utils/config"
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
		a := &Worker1{ID1: uuid.New(), Name1: a.Name1, Description1: a.Description1, FloatNum1: OperationA(a.FloatNum1), CreatedAt1: date_utils.GetNowString()}
		b := &Worker2{ID2: uuid.New(), Name2: b.Name2, Description2: b.Description2, FloatNum2: OperationA(b.FloatNum2), CreatedAt2: date_utils.GetNowString()}
		c := &Worker3{ID3: uuid.New(), Name3: c.Name3, Description3: c.Description3, FloatNum3: OperationA(c.FloatNum3), CreatedAt3: date_utils.GetNowString()}
		d := &Worker4{ID4: uuid.New(), Name4: d.Name4, Description4: d.Description4, FloatNum4: OperationA(d.FloatNum4), CreatedAt4: date_utils.GetNowString()}
		data := config.JsonConfigNodes()

		filea, _ := json.MarshalIndent(a, "", " ")
		os.WriteFile(data.ConfigPath.Workera, filea, 0777) //0644

		fileb, _ := json.MarshalIndent(b, "", " ")
		os.WriteFile(data.ConfigPath.Workerb, fileb, 0777) //0644

		filec, _ := json.MarshalIndent(c, "", " ")
		os.WriteFile(data.ConfigPath.Workerc, filec, 0777) //0644

		filed, _ := json.MarshalIndent(d, "", " ")
		os.WriteFile(data.ConfigPath.Workerd, filed, 0777) //0644

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
