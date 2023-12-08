//CHATGPT PROMPT
/*
1)
Can you write a builtin shell code for the command "ls" in go programming language? I would like it to based off the idea of this code for the "env" command: package builtins

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

2)
that code looks great! however, can you make the first parameter of ListFiles use w as an io.Writer
*/

package builtins

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
