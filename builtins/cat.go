//CHATGPT PROMPT
/*

This code works! Can you make a similar function creating the shell builtin command "cat" in go programming language using this code as a basis: package builtins

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
	"os"
)

func Cat(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: cat [file1 file2 ...]")
	}

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(w, file)
		if err != nil {
			return err
		}
	}

	return nil
}
