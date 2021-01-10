package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type idsFlag []string

func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("My name is: %s, and I am %s", p.name, p.born.String())
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	customValue()
	//flagSets()
}

func customValue() {
	// ./gocliapp --id=1 --id 3 --id 18 -name="John"
	var ids idsFlag
	var p person
	flag.Var(&ids, "id", "the id to be appended to the list")
	flag.Var(&p, "name", "the name of the person")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
}

func flagSets() {
	// ./gocliapp greet --msg=petr
	// ./gocliapp greet --help

	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI APP", "the message for greet command")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("hello and welcome: %s\n", *msgFlag)
	case "help":
		fmt.Println("some help")
	default:
		fmt.Printf("unknown command: %s\n", cmd)

	}
}

func basic() {

	// ./gocliapp greet --msg=petr

	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "CLI APP"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("hello and welcome: %s\n", msg)
	case "help":
		fmt.Println("some help")
	default:
		fmt.Printf("unknown command: %s\n", cmd)

	}
}
