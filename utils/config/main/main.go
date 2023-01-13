package main

import (
	"fmt"

	"github.com/projects/loans/utils/config"
)

func main() {
	data := config.JsonConfigNodes()
	fmt.Println(data.ConfigPath.Workera)
	fmt.Println(data.ConfigDB)
}
