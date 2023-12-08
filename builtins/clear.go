//CHATGPT PROMPT
/*

This code works! Can you make a similar function creating the shell builtin command "clear" in go programming language using this code as a basis: package builtins

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
	//"fmt"
	"io"
	//"os"
	"os/exec"
	"runtime"
)

func ClearScreen(w io.Writer, args ...string) error {
	cmd := getClearCommand()
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

func getClearCommand() *exec.Cmd {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	return cmd
}
