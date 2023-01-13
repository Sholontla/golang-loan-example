package fanout

import "fmt"

func (a *Worker1) processJob1(data string, done chan *Worker1) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, a.Name)
		fmt.Println("Working on data 2", data, a.Description)
		fmt.Println("Working on FloatNum 1", data, a.FloatNum)
		fmt.Printf("\n")
		done <- a
	}()
}

func (b *Worker2) processJob2(data string, done chan *Worker2) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, b.Name)
		fmt.Println("Working on data 2", data, b.Description)
		fmt.Println("Working on FloatNum 1", data, b.FloatNum)
		fmt.Printf("\n")
		done <- b
	}()
}

func (c *Worker3) processJob3(data string, done chan *Worker3) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, c.Name)
		fmt.Println("Working on data 2", data, c.Description)
		fmt.Println("Working on FloatNum 1", data, c.FloatNum)
		fmt.Printf("\n")
		done <- c
	}()
}

func (d *Worker4) processJob4(data string, done chan *Worker4) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, d.Name)
		fmt.Println("Working on data 2", data, d.Description)
		fmt.Println("Working on FloatNum 1", data, d.FloatNum)
		fmt.Printf("\n")
		done <- d
	}()
}
