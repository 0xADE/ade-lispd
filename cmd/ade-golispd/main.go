package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/steelseries/golisp"
)

var loadedScripts = make(chan *golisp.Data)

func main() {
	fmt.Printf("ADE GoLisp scripting\n")

	// Add flag "l" for providing alternative load path instead of default loadPath
	var cmd string
	var loadPath string
	flag.StringVar(&cmd, "e", "", "Run Lisp command")
	flag.StringVar(&loadPath, "l", "~/.config/ade/golisp", "Alternative load path for scripts")
	flag.Parse()

	if loadPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		loadPath = filepath.Join(home, ".config", "ade", "golisp")
	}

	// FIXME load logic:
	// - when loadPath is dir load all lisp files in dir
	// - when loadPath is file load it
	// - when many files available try load init.lsp first
	// - share this logic with zygo
	go loadLispCode(loadPath)
	for r := range loadedScripts {
		fmt.Printf("script loaded: %s\n", golisp.String(r))
	}
	if cmd != "" {
		lispExpr, err := golisp.Parse(cmd)
		if err != nil {
			fmt.Printf("command parsing error: %s", err)
			os.Exit(1)
		}
		res, err := golisp.Eval(lispExpr, golisp.Global)
		if err != nil {
			fmt.Printf("command execution error: %s", err)
			os.Exit(1)
		}
		fmt.Print(golisp.String(res))
	}
}

func loadLispFile(filename string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		// no subpackages
		return nil
	}
	match, _ := filepath.Match("*.lsp", fileInfo.Name())
	if match {
		data, err := golisp.ProcessFile(filename)
		if err != nil {
			return errors.New(fmt.Sprintf("Error GoLisp file: '%s': %s", fileInfo.Name(), err))
		}
		loadedScripts <- data
	}
	return nil
}

// FIXME should load init.lsp first then others
func loadLispCode(path string) {
	err := filepath.Walk(path, loadLispFile)
	if err != nil {
		panic(err)
	}
	close(loadedScripts)
}
