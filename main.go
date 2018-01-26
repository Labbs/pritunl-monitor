package main

import "github.com/Labbs/pritunl-monitor/prometheus"

func main() {
	err := prometheus.Start()
	if err != nil {
		panic(err)
	}
}
