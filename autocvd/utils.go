package autocvd

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

func GetLeastUsedOrFree(leastUsed bool) string {
	if leastUsed {
		return "least-used"
	} else {
		return "free"
	}
}

func AdjustNumGpus(numGpus int, numInstalledGpus int) int {
	if numGpus < 1 || numGpus > numInstalledGpus {
		if numGpus < 1 {
			numGpus = 1
		} else {
			numGpus = numInstalledGpus
		}
		log.Printf(
			"Parameter `numGpus` must be between 1 and %d, setting to %d",
			numInstalledGpus, numGpus,
		)
	}
	return numGpus
}

func GetGpus(numInstalledGpus int) []int {
	gpus := []int{}
	for gpuNum := 0; gpuNum < numInstalledGpus; gpuNum++ {
		gpus = append(gpus, gpuNum)
	}
	return gpus
}

func GetFreeGpus(gpus []int) ([]int, error) {
	freeGpus := []int{}
	for _, gpu := range gpus {
		isFree, err := GetIsFree(gpu)
		if err != nil {
			return nil, err
		}
		if isFree {
			freeGpus = append(freeGpus, gpu)
		}
	}
	return freeGpus, nil
}

func GetAvailableGpusLeastUsed(gpus []int) ([]int, error) {
	availableGpus := []int{}
	freeMemories := map[int]int{}

	for _, gpu := range gpus[:len(gpus)-1] {
		freeGpusMemory, err := GetFreeGpuMemory(gpu)
		if err != nil {
			return nil, err
		}
		freeMemories[gpu] = freeGpusMemory
	}
	for gpu := range freeMemories {
		availableGpus = append(availableGpus, gpu)
	}

	sort.Sort(sort.IntSlice(availableGpus))

	return availableGpus, nil
}

func getAvailableGpus(gpus []int, numGpus int, interval int, progress bool, timeout int, s *spinner.Spinner) ([]int, error) {
	start := time.Now()

	var availableGpus []int
	for {
		freeGpus, err := GetFreeGpus(gpus)
		if err != nil {
			return nil, err
		}
		if len(freeGpus) >= numGpus {
			availableGpus = freeGpus
			break
		}

		for i := 0; i < interval; i++ {
			timePassed := time.Since(start).Seconds()
			if timeout > 0 && timePassed > float64(timeout) {
				return nil, fmt.Errorf("Could not acquire %d GPU(s) before timeout.", numGpus)
			}
			if progress {
				var timeStr string
				if timeout > 0 {
					timeStr = fmt.Sprintf("for %f: > %ds", float64(timeout)-timePassed, len(strconv.Itoa(timeout)))
				} else {
					timeStr = "indefinitely"
				}
				msg := fmt.Sprintf(" %d / %d GPU(s) available (waiting %s, querying every %ds).", len(freeGpus), numGpus, timeStr, interval)
				s.Suffix = msg
			}
			time.Sleep(time.Second * 1)
		}
	}
	return availableGpus, nil
}

func GetAvailableGpusFree(gpus []int, numGpus int, interval int, progress bool, timeout int) ([]int, error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer func() { s.Stop() }()

	availableGpus, err := getAvailableGpus(gpus, numGpus, interval, progress, timeout, s)

	if err != nil {
		return nil, err
	}
	return availableGpus, nil
}

func GetAvailableGpus(leastUsed bool, gpus []int, numGpus int, interval int, progress bool, timeout int) ([]int, error) {
	if leastUsed {
		return GetAvailableGpusLeastUsed(gpus)
	} else {
		return GetAvailableGpusFree(gpus, numGpus, interval, progress, timeout)
	}
}
