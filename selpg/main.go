package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	s        int
	e        int
	l        int
	f        bool
	d        string
	filepath string
)

func captureArgs() {
	flag.IntVar(&s, "s", 0, "specify the start line number to selpg(default 0)")
	flag.IntVar(&e, "e", 0, "specify the end line number to selpg(default 0)")
	flag.IntVar(&l, "l", 72, "specify the number of lines per page(default 72)")
	flag.BoolVar(&f, "f", false, "specify whether use character \\f to do paging")
	flag.StringVar(&d, "d", "", "specify the printer")
	if flag.NArg() > 0 {
		filepath = flag.Arg(0)
	}
	flag.Parse()
}

func validateArgs() bool {
	if e == 0 {
		fmt.Printf("must specify -e")
		return false
	}
	if s > e {
		fmt.Printf("-s must <= -e")
		return false
	}
	if s < 1 || e < 1 {
		fmt.Printf("-s or -e must be positive")
		return false
	}
	if l <= 0 {
		fmt.Printf("-l must be positive")
	}
	if f && l != 72 {
		fmt.Printf("-f and -l can not be both provided")
	}
	return true
}

func readTextFromFile() string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	text := string(b)
	return text
}

func main() {
	captureArgs()
	validateArgs()
	if filepath != "" {
		text := readTextFromFile()
		if f { // -f
			pages := strings.Split(text, "\f")
			for i := s; i != e; i++ {
				fmt.Fprintf(os.Stdout, "%s\f", pages[i])
			}
		} else {
			lines := strings.Split(text, "\n")
			sLines := (s - 1) * l
			eLines := e * l
			for i := sLines; i != eLines; i++ {
				fmt.Fprintf(os.Stdout, "%s\n", lines[i])
			}
		}
	}
	// text := GetTextFromFile("/Users/painterdrown/test")
	// lines := strings.Split(text, "\n")
}
