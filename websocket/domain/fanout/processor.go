package fanout

type Processor struct {
	jobChannel chan string
	done1      chan *Worker1
	done2      chan *Worker2
	done3      chan *Worker3
	done4      chan *Worker4
	workers1   []*Worker1
	workers2   []*Worker2
	workers3   []*Worker3
	workers4   []*Worker4
}
