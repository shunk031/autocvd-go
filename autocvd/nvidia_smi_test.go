package autocvd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

var testCase string

func TestGetInstalledGpus(t *testing.T) {
	testCase = "TestGetInstalledGpus"

	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	numInstalledGpus, err := GetInstalledGpus()
	if err != nil {
		log.Fatal(err)
	}
	if numInstalledGpus != 8 {
		t.Fatal("test failed.")
	}
}

func TestGetIsFree(t *testing.T) {
	testCase = "TestGetIsFree_free"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	isFree, err := GetIsFree(0)
	if err != nil {
		log.Fatal(err)
	}
	if !isFree {
		t.Fatal("test failed")
	}

	testCase = "TestGetIsFree_used"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	isFree, err = GetIsFree(1)
	if err != nil {
		log.Fatal(err)
	}
	if isFree {
		t.Fatal("test failed")
	}
}

func TestGetFreeGpuMemory(t *testing.T) {
	testCase = "TestGetFreeGpuMemory_used"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	mem, err := GetFreeGpuMemory(0)
	if err != nil {
		log.Fatal(err)
	}
	if mem != 6416 {
		t.Fatal("test failed")
	}

	testCase = "TestGetFreeGpuMemory_free"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	mem, err = GetFreeGpuMemory(1)
	if err != nil {
		log.Fatal(err)
	}
	if mem != 48684 {
		t.Fatal("test failed")
	}
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	tc := "TEST_CASE=" + testCase
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", tc}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}
	switch os.Getenv("TEST_CASE") {
	case "TestGetInstalledGpus":
		fmt.Fprint(os.Stdout, "GPU 0: NVIDIA RTX A6000 (UUID: GPU-d22764bb-064f-8d9d-d692-692bf092bd64)\nGPU 1: NVIDIA RTX A6000 (UUID: GPU-6b7fd2f7-5d8f-1a1f-52f7-c3085c029268)\nGPU 2: NVIDIA RTX A6000 (UUID: GPU-28607e5f-56ff-6e5d-f58e-e28989b40db9)\nGPU 3: NVIDIA RTX A6000 (UUID: GPU-30fee18b-2082-ff77-3276-97383558e2ce)\nGPU 4: NVIDIA RTX A6000 (UUID: GPU-c2f4f378-9774-15a5-2a90-90d31ff1602d)\nGPU 5: NVIDIA RTX A6000 (UUID: GPU-0ba6413e-eb52-a509-b2be-abf255d16302)\nGPU 6: NVIDIA RTX A6000 (UUID: GPU-692ff38b-8ea6-902f-818c-832d2190e574)\nGPU 7: NVIDIA RTX A6000 (UUID: GPU-970b83ee-30f9-cd7c-0f2e-11fe70c9d331)\n")
	case "TestGetIsFree_free":
		fmt.Fprint(os.Stdout, "")
	case "TestGetIsFree_used":
		fmt.Fprint(os.Stdout, "3845974")
	case "TestGetFreeGpuMemory_free":
		fmt.Fprint(os.Stdout, "48684")
	case "TestGetFreeGpuMemory_used":
		fmt.Fprint(os.Stdout, "6416")
	case "TestGetFreeGpus_free":
		fmt.Fprint(os.Stdout, "")
	case "TestGetFreeGpus_used":
		fmt.Fprint(os.Stdout, "3845974")
	}
}
