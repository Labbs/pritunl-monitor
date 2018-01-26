package main

import (
	"github.com/pritunl/pritunl-monitor/prometheus"
)

func main() {
	err := prometheus.Start()
	if err != nil {
		panic(err)
	}
}
