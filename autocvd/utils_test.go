package autocvd

import (
	"log"
	"os/exec"
	"testing"
)

func TestGetLeastUsedOrFree(t *testing.T) {
	var leastUsedOrFree string

	leastUsedOrFree = GetLeastUsedOrFree(true)
	if leastUsedOrFree != "least-used" {
		t.Fatal("test failed")
	}

	leastUsedOrFree = GetLeastUsedOrFree(false)
	if leastUsedOrFree != "free" {
		t.Fatal("test failed")
	}
}

func TestAdjustNumGpus(t *testing.T) {
	var (
		numGpus          int
		numInstalledGpus int
	)

	numGpus = 0
	numInstalledGpus = 1
	numGpus = AdjustNumGpus(numGpus, numInstalledGpus)
	if numGpus != 1 {
		t.Fatal("test failed")
	}

	numGpus = 5
	numInstalledGpus = 4
	numGpus = AdjustNumGpus(numGpus, numInstalledGpus)
	if numGpus != 4 {
		t.Fatal("test failed")
	}
}

func TestGetGpus(t *testing.T) {
	var numInstalledGpus int

	numInstalledGpus = 5
	gpus := GetGpus(numInstalledGpus)
	if len(gpus) != numInstalledGpus {
		t.Fatal("test failed")
	}
}

func TestGetFreeGpus(t *testing.T) {
	testCase = "TestGetFreeGpus_free"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	numInstalledGpus := 5
	gpus := GetGpus(numInstalledGpus)
	gpus, err := GetFreeGpus(gpus)
	if err != nil {
		log.Fatal(err)
	}
	if len(gpus) != 5 {
		t.Fatal("test failed")
	}

	testCase = "TestGetFreeGpus_used"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	gpus = GetGpus(numInstalledGpus)
	gpus, err = GetFreeGpus(gpus)
	if err != nil {
		log.Fatal(err)
	}
	if len(gpus) != 0 {
		t.Fatal("test failed")
	}
}
