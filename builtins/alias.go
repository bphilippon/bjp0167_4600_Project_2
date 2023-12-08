//CHATGPT PROMPT
/*

This code also looks great! Can you make a similar function creating the shell builtin command "alias" in go programming language using this code as a basis:  package builtins

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
	//"strings"
)

var aliases map[string]string

func init() {
	aliases = make(map[string]string)
}

func Alias(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// Print existing aliases
		for key, value := range aliases {
			fmt.Fprintf(w, "%s='%s'\n", key, value)
		}
		return nil
	}

	// Set new alias
	if len(args) == 2 {
		aliases[args[0]] = args[1]
		return nil
	}

	return fmt.Errorf("Usage: alias [name[=value] ...]")
}
