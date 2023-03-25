package main

import (
	"flag"
	"fmt"
	"github/com/shunk031/autocvd/autocvd"
	"io/ioutil"
	"log"
	"strings"
)

var (
	numGpus   int
	leastUsed bool
	interval  int
	isExport  bool
	isIdOnly  bool
	isQuiet   bool
	timeout   int
)

func parseArgs() {
	flag.IntVar(&numGpus, "num-gpus", 1, "Number of required GPU.s Defaults to 1.")
	flag.BoolVar(&leastUsed, "least-used", false, "Select least-used GPUs instead of waiting for free GPUs. Defaults to false.")
	flag.IntVar(&interval, "interval", 30, "Interval to query GPU status in seconds. Defaults to 30.")
	flag.BoolVar(&isExport, "export", false, "Add `export` statements such that environment can be sourced.")
	flag.BoolVar(&isIdOnly, "id-only", false, "Return comma-separated GPU IDs only instead of environment variable assignment.")
	flag.BoolVar(&isQuiet, "quiet", false, "Do not print any messages. Defaults to false.")
	flag.IntVar(&timeout, "timeout", -1, "Timeout for waiting in seconds. Defaults to no timeout (-1).")

	flag.Parse()
}

func main() {
	parseArgs()

	if isQuiet {
		log.SetOutput(ioutil.Discard)
	}

	gpus, err := autocvd.Autocvd(numGpus, leastUsed, interval, false, !isQuiet, timeout)
	if err != nil {
		log.Fatal(err)
	}
	gpuDeviceIdStr := strings.Join(gpus, ",")
	if isIdOnly {
		fmt.Println(gpuDeviceIdStr)
	} else {
		var prefix string
		if isExport {
			prefix = "export "
		} else {
			prefix = ""
		}
		fmt.Printf("%sCUDA_DEVICE_ORDER=PCI_BUS_ID %sCUDA_VISIBLE_DEVICES=%s\n", prefix, prefix, gpuDeviceIdStr)
	}
}
