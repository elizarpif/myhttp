package main

import (
	"flag"
	"myhttp/internal/myhttp"
)

func getParallelReqsCount() int {
	var parallelReqsCount int

	flag.IntVar(&parallelReqsCount, "parallel", 10, "max number of parallel requests")
	flag.Parse()

	return parallelReqsCount
}

func getAddresses() []string{
	return flag.Args()
}

func main() {
	count := getParallelReqsCount()
	addresses := getAddresses()

	reqMaker := myhttp.NewRequestsMaker(addresses, count)
	reqMaker.Run()
}
