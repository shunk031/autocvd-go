package autocvd

import (
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Autocvd(numGpus int, leastUsed bool, interval int, setEnv bool, progress bool, timeout int) ([]string, error) {
	numInstalledGpus, err := GetInstalledGpus()
	if err != nil {
		return nil, err
	}
	if numInstalledGpus == 0 {
		return nil, errors.New("There are no GPUs installed.")
	}

	numGpus = AdjustNumGpus(numGpus, numInstalledGpus)
	gpus := GetGpus(numInstalledGpus)

	leastUsedOrFree := GetLeastUsedOrFree(leastUsed)
	log.Printf("Selecting %d %s GPU(s).", numGpus, leastUsedOrFree)

	availableGpus, err := GetAvailableGpus(leastUsed, gpus, numGpus, interval, progress, timeout)
	if err != nil {
		return nil, err
	}

	sort.Sort(sort.IntSlice(availableGpus))
	availableGpus = availableGpus[:numGpus]
	selectedGpus := []string{}
	for _, gpu := range availableGpus {
		selectedGpus = append(selectedGpus, strconv.Itoa(gpu))
	}
	gpuDeviceIdStr := strings.Join(selectedGpus, ",")

	if setEnv {
		os.Setenv("CUDA_DEVICE_ORDER", "PCI_BUS_ID")
		os.Setenv("CUDA_VISIBLE_DEVICES", gpuDeviceIdStr)
		log.Printf("Set environment variables CUDA_VISIBLE)DEVICES=%s and CUDA_DEVICE_ORDER=PCI_BUS_ID.", gpuDeviceIdStr)
	}
	return selectedGpus, nil
}
