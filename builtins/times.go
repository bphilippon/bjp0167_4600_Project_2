//CHATGPT PROMPT
/*

Can you make a function creating the shell builtin command "times" in go programming language using this code as a basis: package builtins

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
	"time"
)

func Times(w io.Writer, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Usage: times takes no arguments")
	}

	// Get the user and system time
	userTime := time.Now().Sub(startTime) - systemTime
	systemTime = time.Now().Sub(startTime)

	// Print the times
	fmt.Fprintf(w, "user\t%s\n", userTime)
	fmt.Fprintf(w, "sys\t%s\n", systemTime)

	return nil
}

var startTime = time.Now()
var systemTime time.Duration
