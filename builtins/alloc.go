package builtins

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

func Alloc(w io.Writer, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Usage: alloc takes no arguments")
	}

	// Allocate memory
	startTime := time.Now()
	allocateMemory()
	allocTime := time.Since(startTime)

	// Print allocation time and memory stats
	fmt.Fprintf(w, "alloc\t%s\n", allocTime)
	printMemoryStats(w)

	return nil
}

func allocateMemory() {
	// Allocate memory (1 MB)
	const oneMB = 1 << 20
	_ = make([]byte, oneMB)
}

func printMemoryStats(w io.Writer) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Fprintf(w, "memory\t%d\n", memStats.Alloc)
	fmt.Fprintf(w, "heap\t%d\n", memStats.HeapAlloc)
	fmt.Fprintf(w, "sys\t%d\n", memStats.Sys)
}
