package main

import (
	"fmt"
	"os"

	"github.com/NVIDIA/go-gpuallocator/gpuallocator"
)

func main() {
	// allocator, err := gpuallocator.NewSimpleAllocator()
	allocator, err := gpuallocator.NewBestEffortAllocator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	for _, gpu := range allocator.GPUs {
		fmt.Printf("%v\n", gpu.Details())
	}

	fmt.Printf("\n")

	for _, i := range []int{1, 2, 4, 8} {
		gpus := allocator.Allocate(i)
		fmt.Printf("Simple allocation of %v GPUs: %v\n", i, gpus)
		allocator.Free(gpus...)
	}
}
