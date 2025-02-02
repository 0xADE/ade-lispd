package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/glycerine/zygomys/zygo"
)

func main() {
	fmt.Printf("ADE Zygomys Lisp scripting\n")

	var (
		cmd      string
		loadPath string
	)
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	loadPath = filepath.Join(home, ".config", "ade", "zygo", "init.zy")

	flag.StringVar(&cmd, "e", "", "Run command")
	flag.StringVar(&loadPath, "l", loadPath, "Alternative load path for scripts")
	flag.Parse()

	lisp := zygo.NewZlisp()

	file, err := os.Open(loadPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = lisp.LoadFile(file)
	if err != nil {
		panic(err)
	}

	if cmd != "" {
		err = lisp.LoadString(cmd)
		if err != nil {
			panic(err)
		}
	}

	expr, err := lisp.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("lisp: %v\n", expr.SexpString(nil))
}
