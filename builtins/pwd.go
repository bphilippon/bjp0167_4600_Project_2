//CHATGPT PROMPT
/*

Great! That code works fine. Can you make a similar function creating the shell builtin command "pwd" in go programming language using this code as a basis: package builtins

import (
	"fmt"
	"io"
	"io/ioutil"
)

func ListFiles(w io.Writer, args ...string) error {
	dir := "."

	if len(args) > 0 {
		dir = args[0]
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}

	return nil
}

*/

package builtins

import (
	"fmt"
	"io"
	"os"
)

func PrintWorkingDirectory(w io.Writer, args ...string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprintln(w, dir)

	return nil
}
