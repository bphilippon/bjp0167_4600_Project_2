//CHATGPT PROMPT
/*

That code looks great! Can you make a similar function creating the shell builtin command "alloc" in go programming language using this code as a basis: package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func EnvironmentVariables(w io.Writer, args ...string) error {
	toRemove := make([]string, 0)
	for i := 0; i < len(args); i++ {
		if args[i] == "-u" {
			if len(args) < i+2 {
				return fmt.Errorf("%w: -u requires an argument", ErrInvalidArgCount)
			}
			toRemove = append(toRemove, args[i+1])
			i++
		}
	}

	toShow := make([]string, 0)
	for _, env := range os.Environ() {
		show := true
		for _, v := range toRemove {
			if strings.HasPrefix(env, v+"=") {
				show = false
				break
			}
		}
		if show {
			toShow = append(toShow, env)
		}
	}

	_, err := fmt.Fprintln(w, strings.Join(toShow, "\n"))

	return err
}

*/

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
