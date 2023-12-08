//CHATGPT PROMPT
/*

This code works! Can you make a similar function creating the shell builtin command "echo" in go programming language using this code as a basis: package builtins

import (
	"fmt"
	"io"
	"strings"
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

*/

package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	fmt.Fprintln(w, strings.Join(args, " "))
	return nil
}
