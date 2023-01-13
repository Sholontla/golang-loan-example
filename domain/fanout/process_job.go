package fanout

func (a *Worker1) processJob1(data string, done chan *Worker1) {
	// Use the data and process the job
	go func() {
		done <- a
	}()
}

func (b *Worker2) processJob2(data string, done chan *Worker2) {
	// Use the data and process the job
	go func() {
		done <- b
	}()
}

func (c *Worker3) processJob3(data string, done chan *Worker3) {
	// Use the data and process the job
	go func() {
		done <- c
	}()
}

func (d *Worker4) processJob4(data string, done chan *Worker4) {
	// Use the data and process the job
	go func() {
		done <- d
	}()
}
