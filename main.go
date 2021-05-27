package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printHelp() {
	fmt.Println("This CLI provides some basic commands to modify files in the current context.")
	fmt.Println("")
	fmt.Println("Following commands are available:")
	fmt.Println("   copyTo <filename>   copies data from stdin to the given filename. When EOF is received the process is stopped.")
	fmt.Println("   mkdir <name>        create a new directory with all sub directories")
	fmt.Println("   help                displayes this message")
	fmt.Println("")
}

func copy(args []string) {
	if len(args) != 1 {
		fmt.Println("ERROR: Exactly one filename is needed")
		printHelp()
		os.Exit(2)
	}
	f, err := os.Create(args[0])
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(f)
	check(err)
	defer f.Close()

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			return
		}
		w.WriteByte(b)
		w.Flush()
	}
}

func mkdir(args []string) {
	if len(args) != 1 {
		fmt.Println("ERROR: Exactly one directory name is needed")
		printHelp()
		os.Exit(2)
	}
	err := os.MkdirAll(args[0], os.ModePerm)
	check(err)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("ERROR: At least one argument is needed\n")
		printHelp()
		os.Exit(1)
	}
	cmd := args[0]
	switch cmd {
	case "copyTo":
		copy(args[1:])
		break
	case "mkdir":
		mkdir(args[1:])
		break
	case "help":
		printHelp()
		break
	default:
		fmt.Printf("ERROR: Unknown command: %s\n", cmd)
		printHelp()
		os.Exit(0)
	}
}
